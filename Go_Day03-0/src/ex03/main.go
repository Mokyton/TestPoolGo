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
		// problem with elasticsearch
		w.Write([]byte("<h1>Wrong lat and lon</h1>"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := NewData("recommendation", places)
	rawJSON, _ := json.Marshal(data)
	_ = json.Indent(&prettyJSON, rawJSON, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write(prettyJSON.Bytes())
	w.WriteHeader(http.StatusOK)
}

func NewData(name string, places []db.Place) Data {
	return Data{Name: name, Places: places}
}
