package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5"
)

type TestGrants struct {
	Grants struct {
		ID           int    `json:"id"`
		Title        string `json:"title"`
		SourceURL    string `json:"source_url"`
		FilterValues struct {
			CuttingOffCriteria []interface{} `json:"cutting_off_criteria"`
			ProjectDirection   []interface{} `json:"project_direction"`
			Amount             int           `json:"amount"`
			LegalForm          []interface{} `json:"legal_form"`
			Age                int           `json:"age"`
		} `json:"filter_values"`
	} `json:"grants"`
}

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
	http.HandleFunc("/grants", grantsHandler)
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

func grantsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	conn, err := pgx.Connect(context.Background(), "postgres://dbgr:2110@localhost:5432/dbgr")
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM grants")
	if err != nil {
		fmt.Printf("Error with query: %v", err)
	}
	defer rows.Close()

	var grant []TestGrants

	for rows.Next() {
		var grantItem TestGrants

		if err := rows.Scan(&grantItem.Grants.ID, &grantItem.Grants.Title, &grantItem.Grants.SourceURL, &grantItem.Grants.FilterValues.ProjectDirection, &grantItem.Grants.FilterValues.Amount, &grantItem.Grants.FilterValues.LegalForm, &grantItem.Grants.FilterValues.Age, &grantItem.Grants.FilterValues.CuttingOffCriteria); err != nil {
			fmt.Printf("Error scanning row: %v", err)
		}
		grant = append(grant, grantItem)
	}

	json, err := json.Marshal(grant)
	if err != nil {
		fmt.Printf("Error with marshal: %v", err)
	}
	fmt.Println(string(json))

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

}
