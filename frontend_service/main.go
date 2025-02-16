package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

var webServiceURL string

func init() {
	webServiceURL = os.Getenv("API_URL")
	if webServiceURL == "" {
		log.Fatal("WEB_SERVICE_URL is not set")
	}
}

func fetchUsers() ([]User, error) {
	// resp, err := http.Get("http://localhost:8081/users") // this is only for local development
	resp, err := http.Get(webServiceURL) //this is for docker containers
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Debugging: Print raw response body
	body, _ := io.ReadAll(resp.Body)
	// fmt.Println("Raw JSON Response:", string(body))

	var users []User
	err = json.Unmarshal(body, &users) // Decode manually
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	// fmt.Println("Parsed Users:", users) // Debugging output
	return users, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("homePage triggered:", r.Method, r.URL.Path) // Debugging

	//tmpl, err := template.ParseFiles("templates/index.html")
	tmpl, err := template.ParseFiles("/app/templates/index.html")

	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	users, err := fetchUsers()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	// fmt.Println("Passing users to template:", users) // Debugging
	tmpl.Execute(w, users)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r) // Simply return 404 for favicon requests
	})
	fmt.Println("Frontend service running on port 8082")
	http.ListenAndServe(":8082", nil)
}
