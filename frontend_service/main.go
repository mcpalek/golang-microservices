package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func fetchUsers() ([]User, error) {
	resp, err := http.Get("http://web_service:8081/users") // Ensure this is correct
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	fmt.Println("Fetched users:", users) // Debugging output
	return users, err
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	users, err := fetchUsers()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	fmt.Println("Passing users to template:", users) // Debugging
	tmpl.Execute(w, users)                           // Render users in template
}

func main() {
	http.HandleFunc("/", homePage)
	fmt.Println("Frontend service running on port 8082")
	http.ListenAndServe(":8082", nil)
}
