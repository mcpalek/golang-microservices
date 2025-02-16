package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/mcpalek/golang-microservices/configloader"
	// "github.com/mcpalek/golang-microservices/configloader"
	// "github.com/mcpalek/golang-microservices/configloader"
	// _ "github.com/denisenkom/go-mssqldb"
)

// ovde je main funkcija bila
func mains() {
	config, err := configloader.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	connString := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;encrypt=disable",
		config.SQLServer.Host, config.SQLServer.Port, config.SQLServer.User, config.SQLServer.Password)
	//fmt.Println("Loaded config:", config)
	start := time.Now()

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE userDB")
	if err != nil {
		log.Println("Database might already exist:", err)
	}

	_, err = db.Exec("USE userDB; CREATE TABLE Users (ID INT PRIMARY KEY IDENTITY, FirstName NVARCHAR(50), LastName NVARCHAR(50))")
	if err != nil {
		log.Println("Table might already exist:", err)
	}

	_, err = db.Exec("USE userDB; INSERT INTO Users (FirstName, LastName) VALUES ('John', 'Doe'), ('Alice', 'Smith')")
	if err != nil {
		log.Println("Error inserting sample data:", err)
	}

	fmt.Println("Database setup completed.")
	elapsed := time.Since(start)
	fmt.Printf("Execution Time: %v\n", elapsed)

}
