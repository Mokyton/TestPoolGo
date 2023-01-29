package main

import (
	"bytes"
	"encoding/json"
	"ex04/db"
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strings"
	"time"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

type AuthTokenClaim struct {
	*jwt.StandardClaims
	User
}

type Data struct {
	Name   string     `json:"name"`
	Places []db.Place `json:"places"`
}

func main() {
	addr := flag.String("addr", ":8888", "Сетевой адрес HTTP")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/recommend", myMiddleware(handler))
	mux.HandleFunc("/api/get_token", handlerGenerateToken)

	fmt.Printf("Запуск сервера на %s\n", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerGenerateToken(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errHandler(w, "Invalid json", http.StatusInternalServerError)
		return
	}

	expiresAt := time.Now().Add(time.Minute * 1).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &AuthTokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		User{user.Username, user.Password},
	}

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		errHandler(w, "Cannot signed token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Token", tokenString)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: time.Now().Add(time.Minute * 10).Unix(),
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	var prettyJSON bytes.Buffer

	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	places, err := db.New().GetPlaces(3, lat, lon)
	if err != nil {
		errHandler(w, "Invalid parameters error", http.StatusBadRequest)
		return
	}

	data := NewData("recommendation", places)
	rawJSON, _ := json.Marshal(data)
	_ = json.Indent(&prettyJSON, rawJSON, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(prettyJSON.Bytes())
	w.WriteHeader(http.StatusOK)
}

func myMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(bearer, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error")
			}
			return sampleSecretKey, nil
		})

		if err != nil {
			errHandler(w, "Token parsing error", http.StatusInternalServerError)
			return
		}

		if !token.Valid {
			errHandler(w, "Authorization error", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func NewData(name string, places []db.Place) Data {
	return Data{Name: name, Places: places}
}

func errHandler(w http.ResponseWriter, message string, errStatus int) {
	jsn := struct {
		Error string `json:"error"`
	}{message}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(jsn)
	w.WriteHeader(errStatus)
}
