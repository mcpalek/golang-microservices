package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", homePage)
	fmt.Println("Frontend service running on port 8082")
	http.ListenAndServe(":8082", nil)
}
