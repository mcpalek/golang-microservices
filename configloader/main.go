package main

import (
	"fmt"
	"log"

	"github.com/mcpalek/golang-microservices/configloader" // Adjust the path according to your project structure
)

func main() {
	config, err := configloader.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	fmt.Printf("Server: %s\n", config.Server)
	fmt.Printf("Port: %s\n", config.Port)
	fmt.Printf("User: %s\n", config.User)
	fmt.Printf("Password: %s\n", config.Password)
}
