package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var tmpl *template.Template

// init()
func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

// main is the entry point of the
func main() {
	// define gorilla mux
	r := mux.NewRouter()
	// static file server
	fs := http.FileServer(http.Dir("assets/"))
	// http.Handle("/assets/", http.StripPrefix("/assets", fs))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fs))

	//
	r.HandleFunc("/", HomeHandler)

	//
	r.HandleFunc("/about", AboutHandler)

	//
	r.HandleFunc("/contact", ContactsHander)

	//
	r.HandleFunc("/login", LoginHandler)

	//
	r.HandleFunc("/signup", SignupHandler)

	// start server
	http.ListenAndServe(":9090", r)
}

// HomeHandler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.layout.html", nil)
}

// AboutHandler
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.layout.html", nil)
}

// ContactsHander
func ContactsHander(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "contact.layout.html", nil)
}

// LoginHandler
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "login.layout.html", nil)
}

// SignupHandler
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "signup.layout.html", nil)
}
