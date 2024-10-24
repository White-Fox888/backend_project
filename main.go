package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

type Identification struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokenTest struct {
	Token string `json:"token"`
}

var authToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHAiOiIyMDIzLTEyLTE4VDEyOjI5OjE5LjEwNjg0MTQzOVoiLCJVc2VyTG9naW4iOiJhZG1pbiJ9.0Dvg7vFTrdSX2F4751ae6Id9weC5ATvF1sQPuvejiFE"

type Grant struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	SourceURL    string `json:"source_url"`
	FilterValues struct {
		CuttingOffCriteria []int `json:"cutting_off_criteria"`
		ProjectDirection   []int `json:"project_direction"`
		Amount             int   `json:"amount"`
		LegalForm          []int `json:"legal_form"`
		Age                int   `json:"age"`
	} `json:"filter_values"`
}

type FilterMapping struct {
	Age struct {
		Title   string `json:"title"`
		Mapping struct {
		} `json:"mapping"`
	} `json:"age"`
	ProjectDirection struct {
		Title   string `json:"title"`
		Mapping struct {
			Num0 struct {
				Title string `json:"title"`
			} `json:"0"`
			Num1 struct {
				Title string `json:"title"`
			} `json:"1"`
			Num2 struct {
				Title string `json:"title"`
			} `json:"2"`
			Num3 struct {
				Title string `json:"title"`
			} `json:"3"`
		} `json:"mapping"`
	} `json:"project_direction"`
	LegalForm struct {
		Title   string `json:"title"`
		Mapping struct {
			Num0 struct {
				Title string `json:"title"`
			} `json:"0"`
			Num1 struct {
				Title string `json:"title"`
			} `json:"1"`
			Num2 struct {
				Title string `json:"title"`
			} `json:"2"`
		} `json:"mapping"`
	} `json:"legal_form"`
	CuttingOffCriteria struct {
		Title   string `json:"title"`
		Mapping struct {
			Num0 struct {
				Title string `json:"title"`
			} `json:"0"`
			Num1 struct {
				Title string `json:"title"`
			} `json:"1"`
			Num2 struct {
				Title string `json:"title"`
			} `json:"2"`
			Num3 struct {
				Title string `json:"title"`
			} `json:"3"`
		} `json:"mapping"`
	} `json:"cutting_off_criteria"`
	Amount struct {
		Title   string `json:"title"`
		Mapping struct {
		} `json:"mapping"`
	} `json:"amount"`
}

type FilterOrder []string

type MetaPages struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}

type DataFilters struct {
	Data struct {
		ProjectDirection   []int `json:"project_direction"`
		Amount             int   `json:"amount"`
		LegalForm          []int `json:"legal_form"`
		Age                int   `json:"age"`
		CuttingOffCriteria []int `json:"cutting_off_criteria"`
	} `json:"data"`
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/grants", grantsHandler)
	http.HandleFunc("/grants/{id}", grantIDHandler)
	http.HandleFunc("/grants/{id}/filters", filterHandler)
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := pgx.Connect(context.Background(), "postgres://dbgr:2110@localhost:5432/dbgr")
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var ident Identification
	err = json.NewDecoder(r.Body).Decode(&ident)
	if err != nil {
		fmt.Printf("Неверный формат данных: %v", err)
		return
	}

	var hashedPassword []byte
	err = conn.QueryRow(context.Background(), "SELECT password FROM users WHERE login=$1", ident.Login).Scan(&hashedPassword)
	if err != nil {
		fmt.Printf("test: %v", err)
		return
	}

	var isValid bool
	err = conn.QueryRow(context.Background(), "SELECT crypt($1, password) = password FROM users WHERE login = $2", ident.Password, ident.Login).Scan(&isValid)
	if err != nil {
		fmt.Printf("Unauthorized: %v", err)
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
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func grantsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
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

	var Grants []Grant
	var grantItem Grant
	for rows.Next() {
		if err := rows.Scan(&grantItem.ID, &grantItem.Title, &grantItem.SourceURL,
			&grantItem.FilterValues.ProjectDirection, &grantItem.FilterValues.Amount,
			&grantItem.FilterValues.LegalForm, &grantItem.FilterValues.Age,
			&grantItem.FilterValues.CuttingOffCriteria); err != nil {
			fmt.Printf("Error scanning row: %v", err)
		}
		Grants = append(Grants, grantItem)
	}

	var FiltersMapping FilterMapping
	err = conn.QueryRow(context.Background(), "SELECT * FROM filters_mapping").Scan(
		&FiltersMapping.Age,
		&FiltersMapping.ProjectDirection,
		&FiltersMapping.LegalForm,
		&FiltersMapping.CuttingOffCriteria,
		&FiltersMapping.Amount)
	if err != nil {
		fmt.Printf("Error with query & scanning: %v", err)
	}

	var Meta MetaPages
	err = conn.QueryRow(context.Background(), "SELECT * FROM meta").Scan(
		&Meta.CurrentPage, &Meta.TotalPages)
	if err != nil {
		fmt.Printf("Error with query & scanning: %v", err)
	}

	var FiltersOrder FilterOrder

	DataGrants := struct {
		Grants         []Grant       `json:"grants"`
		FiltersMapping FilterMapping `json:"filters_mapping"`
		FiltersOrder   FilterOrder   `json:"filters_order"`
		Meta           MetaPages     `json:"meta"`
	}{
		Grants:         Grants,
		FiltersMapping: FiltersMapping,
		FiltersOrder:   FiltersOrder,
		Meta:           Meta,
	}

	json, err := json.Marshal(DataGrants)
	if err != nil {
		fmt.Printf("Error with marshal: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func grantIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := pgx.Connect(context.Background(), "postgres://dbgr:2110@localhost:5432/dbgr")
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	var GrantID Grant
	ID := r.PathValue("id")
	err = conn.QueryRow(context.Background(), "SELECT * FROM grants WHERE id = $1", ID).Scan(
		&GrantID.ID, &GrantID.Title, &GrantID.SourceURL,
		&GrantID.FilterValues.ProjectDirection, &GrantID.FilterValues.Amount,
		&GrantID.FilterValues.LegalForm, &GrantID.FilterValues.Age,
		&GrantID.FilterValues.CuttingOffCriteria)
	if err != nil {
		fmt.Printf("Error with query & scanning: %v", err)
	}

	var FilterMappingID FilterMapping
	err = conn.QueryRow(context.Background(), "SELECT * FROM filters_mapping").Scan(
		&FilterMappingID.Age,
		&FilterMappingID.ProjectDirection,
		&FilterMappingID.LegalForm,
		&FilterMappingID.CuttingOffCriteria,
		&FilterMappingID.Amount)
	if err != nil {
		fmt.Printf("Error with query & scanning: %v", err)
	}

	var FiltersOrder FilterOrder

	DataGrantID := struct {
		GrantID         Grant         `json:"grant"`
		FilterMappingID FilterMapping `json:"filters_mapping"`
		FiltersOrder    FilterOrder   `json:"filters_order"`
	}{
		GrantID:         GrantID,
		FilterMappingID: FilterMappingID,
		FiltersOrder:    FiltersOrder,
	}
	json, err := json.Marshal(DataGrantID)
	if err != nil {
		fmt.Printf("Error with marshal: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := pgx.Connect(context.Background(), "postgres://dbgr:2110@localhost:5432/dbgr")
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	var DataID DataFilters
	err = json.NewDecoder(r.Body).Decode(&DataID)
	if err != nil {
		fmt.Printf("Неверный формат данных: %v", err)
		return
	}

	ID := r.PathValue("id")
	intID, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Printf("Error with converting: %v", err)
	}

	request, err := conn.Exec(context.Background(), "UPDATE grants SET project_directions = $1, amount = $2, legal_forms = $3, age = $4, cutting_off_criterea = $5 WHERE id = $6",
		&DataID.Data.ProjectDirection, &DataID.Data.Amount, &DataID.Data.LegalForm, &DataID.Data.Age, &DataID.Data.CuttingOffCriteria, intID)
	if err != nil {
		fmt.Printf("Error with scanning: %v", err)
	}
	if request.RowsAffected() != 1 {
		fmt.Printf("Error with updating: %v", err)
	}
	w.WriteHeader(http.StatusNoContent)
}
