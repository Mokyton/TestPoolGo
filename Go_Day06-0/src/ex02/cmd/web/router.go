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
	router.HandleFunc("/", app.GlobalMiddleware(app.home))
	router.HandleFunc("/admin", app.GlobalMiddleware(app.signIn))
	router.HandleFunc("/login", app.GlobalMiddleware(app.loginHandler))
	router.HandleFunc("/createThought", app.GlobalMiddleware(app.CreateThought))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))
	return router
}
