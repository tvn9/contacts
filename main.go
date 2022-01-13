package main

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template

// init()
func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

// main is the entry point of the
func main() {

	// static file server
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	// define gorilla mux
	// r := mux.NewRouter()

	//
	http.HandleFunc("/", HomeHandler)

	//
	http.HandleFunc("/about", AboutHandler)

	//
	http.HandleFunc("/contact", ContactsHander)

	//
	http.HandleFunc("/login", LoginHandler)

	//
	http.HandleFunc("/signup", SignupHandler)

	// start server
	http.ListenAndServe(":9090", nil)
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
