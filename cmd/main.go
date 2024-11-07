package main

import (
	"backend_project/config"
	conndb "backend_project/db"
	"backend_project/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/check", handlers.CheckHandler)
	http.HandleFunc("/grants", handlers.GrantsHandler)
	http.HandleFunc("/grants/{id}", handlers.GrantIDHandler)
	http.HandleFunc("/grants/{id}/filters", handlers.FilterHandler)
	http.ListenAndServe(":8080", nil)

	var conf = config.GetEnv()
	var DB = conf.Database

	db, err := conndb.InitDB(DB)
	if err != nil {
		fmt.Printf("Database initialization error: %v/n", err)
		return
	}
	defer db.Close()

	handlers.SetDatabase(db)
}
