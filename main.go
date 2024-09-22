package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Identification struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokenTest struct {
	Token string `json:"token"`
}

var authToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHAiOiIyMDIzLTEyLTE4VDEyOjI5OjE5LjEwNjg0MTQzOVoiLCJVc2VyTG9naW4iOiJhZG1pbiJ9.0Dvg7vFTrdSX2F4751ae6Id9weC5ATvF1sQPuvejiFE"

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/check", checkHandler)
	// http.HandleFunc("/grants", grantsHandler)
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var ident Identification
	err := json.NewDecoder(r.Body).Decode(&ident)
	if err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	if ident.Login == "admin" && ident.Password == "correct_password" {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TokenTest{Token: authToken})
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token := authHeader[len("Bearer "):]
	if token != authToken {
		http.Error(w, "No Content", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// func grantsHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}

// 	authHeader := r.Header.Get("Authorization")
// 	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		return
// 	}
// }
