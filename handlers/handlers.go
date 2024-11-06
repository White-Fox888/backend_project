package handlers

import (
	"backend_project/config"
	"backend_project/structs"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}
}

var conf = config.GetEnv()
var Key = conf.SecretKey
var DB = conf.Database

var err error
var conn *pgx.Conn

// func InitDB() (*pgx.Conn, error) {
// 	conn, err := pgx.Connect(context.Background(), DB)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return conn, nil
// }

// func closeDB(conn *pgx.Conn) {
// 	err := conn.Close(context.Background())
// 	if err != nil {
// 		fmt.Println("Ошибка при закрытии соединения:", err)
// 	}
// }

var FiltersOrder structs.FilterOrder = structs.FilterOrder{"project_direction", "amount", "legal_form", "age", "cutting_off_criteria"}

func GenerateToken(myClaims *structs.Claims) ([]byte, error) {
	var mySigningKey = []byte(Key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	TJWT, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	tokenJson := structs.Token{Token: TJWT}
	json, err := json.Marshal(tokenJson)
	if err != nil {
		fmt.Printf("Error with marshal: %v", err)
	}
	return json, nil
}

func ValidateToken(tokenString string) (bool, error) {
	var mySigningKey = []byte(Key)
	valMethod := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(mySigningKey), nil
	}
	valClaims := &structs.Claims{
		RegisteredClaims: jwt.RegisteredClaims{},
		Login:            "admin",
	}

	parsedToken, err := jwt.ParseWithClaims(tokenString, valClaims, valMethod)
	if err != nil {
		fmt.Printf("Parsing error: %v", err)
	}
	if !parsedToken.Valid {
		fmt.Printf("Invalid token")
	}
	return true, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	conn, err := pgx.Connect(context.Background(), DB)
	if err != nil {
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		fmt.Printf("Unable to connect to database: %v\n", err)
		return
	}
	defer conn.Close(context.Background())

	var ident structs.Identification
	err = json.NewDecoder(r.Body).Decode(&ident)
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		fmt.Printf("Invalid data format: %v", err)
		return
	}

	var isValid bool
	err = conn.QueryRow(context.Background(), "SELECT crypt($1, password) = password FROM users WHERE login = $2", ident.Password, ident.Login).Scan(&isValid)
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		fmt.Printf("Error with query: %v", err)
		return
	}

	if isValid {
		myClaims := &structs.Claims{
			RegisteredClaims: jwt.RegisteredClaims{},
			Login:            ident.Login,
		}

		token, err := GenerateToken(myClaims)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(token)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	val, err := ValidateToken(tokenString)
	if err != nil || !val {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GrantsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	val, err := ValidateToken(tokenString)
	if err != nil || !val {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	conn, err := pgx.Connect(context.Background(), DB)
	if err != nil {
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		fmt.Printf("Unable to connect to database: %v\n", err)
		return
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM grants")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error with query: %v", err)
		return
	}
	defer rows.Close()

	var Grants []structs.Grant
	var grantItem structs.Grant
	for rows.Next() {
		if err := rows.Scan(&grantItem.ID, &grantItem.Title, &grantItem.SourceURL,
			&grantItem.FilterValues.ProjectDirection, &grantItem.FilterValues.Amount,
			&grantItem.FilterValues.LegalForm, &grantItem.FilterValues.Age,
			&grantItem.FilterValues.CuttingOffCriteria); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Printf("Error scanning row: %v", err)
			return
		}
		Grants = append(Grants, grantItem)
	}

	var FiltersMapping structs.FilterMapping
	err = conn.QueryRow(context.Background(), "SELECT * FROM filters_mapping").Scan(
		&FiltersMapping.Age,
		&FiltersMapping.ProjectDirection,
		&FiltersMapping.LegalForm,
		&FiltersMapping.CuttingOffCriteria,
		&FiltersMapping.Amount)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error with query & scanning: %v", err)
		return
	}

	var Meta structs.MetaPages
	err = conn.QueryRow(context.Background(), "SELECT * FROM meta").Scan(
		&Meta.CurrentPage, &Meta.TotalPages)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error with query & scanning: %v", err)
		return
	}

	DataGrants := struct {
		Grants         []structs.Grant       `json:"grants"`
		FiltersMapping structs.FilterMapping `json:"filters_mapping"`
		FiltersOrder   structs.FilterOrder   `json:"filters_order"`
		Meta           structs.MetaPages     `json:"meta"`
	}{
		Grants:         Grants,
		FiltersMapping: FiltersMapping,
		FiltersOrder:   FiltersOrder,
		Meta:           Meta,
	}

	json, err := json.Marshal(DataGrants)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error with marshal: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func GrantIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	val, err := ValidateToken(tokenString)
	if err != nil || !val {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	conn, err := pgx.Connect(context.Background(), DB)
	if err != nil {
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		fmt.Printf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	var GrantID structs.Grant
	ID := r.PathValue("id")
	err = conn.QueryRow(context.Background(), "SELECT * FROM grants WHERE id = $1", ID).Scan(
		&GrantID.ID, &GrantID.Title, &GrantID.SourceURL,
		&GrantID.FilterValues.ProjectDirection, &GrantID.FilterValues.Amount,
		&GrantID.FilterValues.LegalForm, &GrantID.FilterValues.Age,
		&GrantID.FilterValues.CuttingOffCriteria)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error with query & scanning: %v", err)
		return
	}

	var FilterMappingID structs.FilterMapping
	err = conn.QueryRow(context.Background(), "SELECT * FROM filters_mapping").Scan(
		&FilterMappingID.Age,
		&FilterMappingID.ProjectDirection,
		&FilterMappingID.LegalForm,
		&FilterMappingID.CuttingOffCriteria,
		&FilterMappingID.Amount)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error with query & scanning: %v", err)
		return
	}

	DataGrantID := struct {
		GrantID         structs.Grant         `json:"grant"`
		FilterMappingID structs.FilterMapping `json:"filters_mapping"`
		FiltersOrder    structs.FilterOrder   `json:"filters_order"`
	}{
		GrantID:         GrantID,
		FilterMappingID: FilterMappingID,
		FiltersOrder:    FiltersOrder,
	}
	json, err := json.Marshal(DataGrantID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error with marshal: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	val, err := ValidateToken(tokenString)
	if err != nil || !val {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	conn, err := pgx.Connect(context.Background(), DB)
	if err != nil {
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		fmt.Printf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	var DataID structs.DataFilters
	err = json.NewDecoder(r.Body).Decode(&DataID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Invalid data format: %v", err)
		return
	}

	ID := r.PathValue("id")
	intID, err := strconv.Atoi(ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error with converting: %v", err)
		return
	}

	request, err := conn.Exec(context.Background(), "UPDATE grants SET project_directions = $1, amount = $2, legal_forms = $3, age = $4, cutting_off_criterea = $5 WHERE id = $6",
		&DataID.Data.ProjectDirection, &DataID.Data.Amount, &DataID.Data.LegalForm, &DataID.Data.Age, &DataID.Data.CuttingOffCriteria, intID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error with scanning: %v", err)
		return
	}
	if request.RowsAffected() != 1 {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error with updating: %v", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
