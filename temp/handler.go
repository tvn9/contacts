package main

import (
	"log"
	"net/http"
	"text/template"
)

//
var tmpl *template.Template

// init initialize the template variable/type
func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

// Home
func home(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Printf("Fails to execute template %v\n", err)
	}
}

// Contact
func contact(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "contacts.html", nil)
	if err != nil {
		log.Printf("Fails to execute template %v\n", err)
	}
}

// About
func about(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		log.Printf("Fails to execute template %v\n", err)
	}
}

// Login
func login(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Printf("Fails to execute template %v\n", err)
	}
}

// Signup
func signup(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "signup.html", nil)
	if err != nil {
		log.Printf("Fails to execute template %v\n", err)
	}
}
