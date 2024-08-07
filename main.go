package main

import (
	"encoding/json"
	"net/http"
)

type Identification struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
	http.HandleFunc("/admin/api/v1/auth/login", loginHandler)
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var ident Identification
	err := json.NewDecoder(r.Body).Decode(&ident)
	if err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}
	if ident.Login == "admin" && ident.Password == "correct_password" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Успешная авторизация"))
	} else {
		http.Error(w, "Code: 401", http.StatusUnauthorized)
	}
}
