package main

import (
	"backend_project/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/check", handlers.CheckHandler)
	http.HandleFunc("/grants", handlers.GrantsHandler)
	http.HandleFunc("/grants/{id}", handlers.GrantIDHandler)
	http.HandleFunc("/grants/{id}/filters", handlers.FilterHandler)
	http.ListenAndServe(":8080", nil)
}
