package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	Host string
	Port string
}

// main
func main() {

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	// handleFunc
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/contact", contact)

	host := Server{
		Host: "",
		Port: ":8080",
	}

	fmt.Printf("Starting server on port %s\n", host.Port)
	http.ListenAndServe(host.Port, nil)
}
