package main

import (
	"html/template"
	"net/http"
)

func (app *application) adminHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin" {
		http.NotFound(w, r)
		return
	}

	t, err := template.ParseFiles("./ui/html/login.gohtml")

	if err != nil {
		http.Error(w, err.Error(), 555)
		return
	}
	t.Execute(w, nil)

}
func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello world"))

}
