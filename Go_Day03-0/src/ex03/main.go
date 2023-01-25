package main

import (
	"bytes"
	"encoding/json"
	"ex03/db"
	"flag"
	"fmt"
	"net/http"
)

type Data struct {
	Name   string     `json:"name"`
	Places []db.Place `json:"places"`
}

func main() {
	addr := flag.String("addr", ":8888", "Сетевой адрес HTTP")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/recommend", handler)

	fmt.Printf("Запуск сервера на %s\n", *addr)
	http.ListenAndServe(*addr, mux)
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
