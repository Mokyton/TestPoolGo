package main

import (
	"ex01/pkg/db"
	"flag"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

type Data struct {
	Total    int
	Places   []db.Place
	LastPage int
	PrevPage int
	NextPage int
	Rng      []int
}

func main() {
	addr := flag.String("addr", ":9090", "Сетевой адрес HTTP")
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", handler)

	fmt.Printf("Запуск сервера на %s", *addr)
	http.ListenAndServe(*addr, mux)

}

func handler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("page")
	page, err := strconv.Atoi(param)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Invalid 'page' value: '%s'", param)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		w.Write([]byte("Can't parse html"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	places, total, err := db.New().GetPlaces(10, page)
	if err != nil {
		w.Write([]byte("Page didn't found"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data := NewData(places, total, page)
	err = tmpl.Execute(w, data)
	if err != nil {
		w.Write([]byte("Wrong template "))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
}

func NewData(places []db.Place, total int, page int) Data {
	d := Data{Total: total, Places: places}
	d.LastPage = int(math.Ceil(float64(total) / 10.0))
	d.PrevPage = page - 1
	if d.PrevPage < 0 {
		d.PrevPage = 1
	}
	d.NextPage = page + 1
	if d.NextPage > d.LastPage {
		d.NextPage = d.LastPage
	}
	v := d.LastPage - page
	if v > 10 {
		v = 10
	}
	rng := make([]int, v, v)
	for i := 0; i < v; i++ {
		rng[i] = page + i + 1
	}
	d.Rng = rng
	return d
}
