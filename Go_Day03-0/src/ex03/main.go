package main

import (
	"bytes"
	"encoding/json"
	"ex03/db"
	"flag"
	"fmt"
	"net/http"
	"strconv"
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

	latStr := r.URL.Query().Get("lat")
	lonStr := r.URL.Query().Get("lon")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		// cannot get lat
		fmt.Println(err)
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		// cannot get lon
		fmt.Println(err)
	}

	places, err := db.New().GetPlaces(3, lat, lon)
	if err != nil {
		// problem with elasticsearch
		fmt.Println(err)
	}
	data := NewData("recommendation", places)
	rawJSON, err := json.Marshal(data)
	err = json.Indent(&prettyJSON, rawJSON, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write(prettyJSON.Bytes())
	w.WriteHeader(http.StatusOK)
}

func path(path string) string {
	for i := len(path) - 1; i > 0; i-- {
		if string(path[i]) == "/" {
			path = path[i+1:]
			break
		}
	}
	return path
}

func NewData(name string, places []db.Place) Data {
	return Data{Name: name, Places: places}
}
