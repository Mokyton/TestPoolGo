package main

import (
	"ex01/pkg/db"
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

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", handler)

	fmt.Println("Server start")
	http.ListenAndServe(":9090", mux)
	//da := pkg.Store().GetPlaces(10, 20)
	//fmt.Println(d)
	//fmt.Println(test)
}

func handler(w http.ResponseWriter, r *http.Request) {

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		//fmt.Println(r.URL.Query().)
		//fmt.Println(err)
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {

	}
	places, total, err := db.New().GetPlaces(10, page)
	if err != nil {
		fmt.Println(err)
	}
	data := NewData(places, total, page)

	err = tmpl.Execute(w, data)

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
	rng := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		rng[i] = page + i + 1
	}
	d.Rng = rng
	return d
}
