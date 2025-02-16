package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/mcpalek/golang-microservices/configloader"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	config, err := configloader.LoadConfig()
	if err != nil {
		http.Error(w, "Config error", http.StatusInternalServerError)
		fmt.Println("Config error:", err) // Debugging
		return
	}

	connString := fmt.Sprintf("server=%s;port=%s;database=userDB;user id=%s;password=%s;encrypt=disable",
		config.SQLServer.Host, config.SQLServer.Port, config.SQLServer.User, config.SQLServer.Password)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		fmt.Println("Database connection error:", err) // Debugging
		return
	}
	defer db.Close()

	fmt.Println("Database connection successful") // Debugging

	rows, err := db.Query("SELECT id, FirstName, LastName FROM Users")
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		fmt.Println("Query Error:", err) // Debugging
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName); err != nil {
			http.Error(w, "Error scanning users", http.StatusInternalServerError)
			fmt.Println("Error scanning row:", err) // Debugging
			return
		}
		users = append(users, u)
	}

	if len(users) == 0 {
		http.Error(w, "No users found", http.StatusNotFound)
		fmt.Println("No users found in database") // Debugging
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(users)
}

// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	config, err := configloader.LoadConfig()
// 	if err != nil {
// 		http.Error(w, "Config error", http.StatusInternalServerError)
// 		return
// 	}

// 	connString := fmt.Sprintf("server=%s;port=%s;database=userDB;user id=%s;password=%s;encrypt=disable",
// 		config.SQLServer.Host, config.SQLServer.Port, config.SQLServer.User, config.SQLServer.Password)

// 	db, err := sql.Open("sqlserver", connString)
// 	if err != nil {
// 		http.Error(w, "Database connection error", http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT id, FirstName, LastName FROM Users")
// 	if err != nil {
// 		http.Error(w, "Error fetching users", http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	var users []User
// 	for rows.Next() {
// 		var u User
// 		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName); err != nil {
// 			http.Error(w, "Error scanning users", http.StatusInternalServerError)
// 			return
// 		}
// 		users = append(users, u)
// 	}
// 	fmt.Println(rows)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(users)
// }

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("API service running on port 8081")
	http.ListenAndServe(":8081", nil)
}
