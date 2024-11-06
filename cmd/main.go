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

	// conn, err := initDB()
	// if err != nil {
	// 	fmt.Print("Ошибка инициализации базы данных:", err)
	// }
	// defer closeDB(conn)

	// conn, err := pgx.Connect(context.Background(), DB)
	// if err != nil {
	// 	fmt.Printf("Unable to connect to database: %v\n", err)
	// }
	// defer conn.Close(context.Background())
}

// var conf = config.GetEnv()
// var DB = conf.Database

// func initDB() (*pgx.Conn, error) {
// 	conn, err := pgx.Connect(context.Background(), DB)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return conn, nil
// }

// func closeDB(conn *pgx.Conn) {
// 	err := conn.Close(context.Background())
// 	if err != nil {
// 		fmt.Print("Ошибка при закрытии соединения:", err)
// 	}
// }
