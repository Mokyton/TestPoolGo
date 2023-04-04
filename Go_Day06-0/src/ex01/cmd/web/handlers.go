package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.NotFound(w, r)
		return
	}
	r.ParseForm()
	if r.Form.Get("login") == "admin" && r.Form.Get("password") == "admin" {
		ses, err := cookieStore.Get(r, cookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ses.Values[sesKeyLogin] = "admin"
		err = cookieStore.Save(r, w, ses)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func (app *application) signIn(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./ui/html/signin.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.Execute(w, nil)
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	ses, err := cookieStore.Get(r, cookieName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	login, ok := ses.Values[sesKeyLogin].(string)
	if !ok {
		login = "anonymous"
	}

	var page, nextPage, prevPage int
	var limit = 3

	if strings.HasPrefix(r.URL.RawQuery, "page=") {
		num, err := strconv.Atoi(strings.TrimPrefix(r.URL.RawQuery, "page="))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		page = num
	} else {
		page = 1
	}

	total := app.thoughts.GetCountOfThoughts()
	thoughts, _ := app.thoughts.GetThoughts(limit, limit*(page-1))
	if page*limit < total {
		nextPage = page + 1
	} else {
		nextPage = 0
	}
	prevPage = page - 1

	user, ok := users[login]
	if !ok {
		user = User{Name: login, FullAccess: false}
	}
	user.Thoughts = thoughts
	user.Previous = prevPage
	user.Next = nextPage
	t, err := template.ParseFiles("./ui/html/home.gohtml")
	t.Execute(w, user)
}

func (app *application) CreateThought(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/createThought" || r.Method != "POST" {
		http.Error(w, errors.New("Bad request ").Error(), http.StatusBadRequest)
		return
	}
	r.ParseForm()
	fmt.Println(r.Form.Get("Title"), r.Form.Get("Content"))
	app.thoughts.Insert(r.Form.Get("Title"), r.Form.Get("Content"))
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
