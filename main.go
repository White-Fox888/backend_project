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
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
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
		w.Write([]byte("Успешная авторизация"))
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
