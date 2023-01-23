package main

import (
	"bytes"
	"encoding/json"
	"ex02/db"
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

type Data struct {
	Name   string     `json:"name"`
	Total  int        `json:"total"`
	Places []db.Place `json:"places"`
}

func main() {
	addr := flag.String("addr", ":8888", "Сетевой адрес HTTP")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/places/", handler)

	fmt.Printf("Запуск сервера на %s", *addr)
	http.ListenAndServe(*addr, mux)

}

func handler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("page")
	page, err := strconv.Atoi(param)
	if err != nil {
		if param == "" {
			errWraper(w, http.StatusNotFound, "Not Found")
		} else {
			errWraper(w, http.StatusBadRequest, fmt.Sprintf("Invalid 'page' value: '%s'", param))
		}
		return
	}

	places, total, err := db.New().GetPlaces(10, page)
	if err != nil {
		errWraper(w, http.StatusNotFound, "Page doesn't exist")
		return
	}

	data := NewData(places, total, "places")
	w.Header().Set("Content-Type", "application/json")
	rawJSON, err := json.Marshal(data)
	if err != nil {
		errWraper(w, http.StatusInternalServerError, "Cannot marshal json")
		return
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, rawJSON, "", "\t")
	w.Write(prettyJSON.Bytes())
	w.WriteHeader(200)
}

func NewData(places []db.Place, total int, index string) Data {
	return Data{Name: index, Total: total, Places: places}
}

func errWraper(w http.ResponseWriter, status int, err string) {
	e := struct {
		Error string `json:"error"`
	}{Error: err}
	rawJSON, _ := json.Marshal(e)
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, rawJSON, "", "\t")
	w.Write(prettyJSON.Bytes())
	w.WriteHeader(status)
}
