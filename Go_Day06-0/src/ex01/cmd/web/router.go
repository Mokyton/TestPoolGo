package main

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var cookieStore = sessions.NewCookieStore([]byte("secret"))

const cookieName = "MyCookie"

type sesKey int

const (
	sesKeyLogin sesKey = iota
)

func (app *application) Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", app.home)
	router.HandleFunc("/admin", app.signIn)
	router.HandleFunc("/login", app.loginHandler)
	router.HandleFunc("/createThought", app.CreateThought)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))
	return router
}
