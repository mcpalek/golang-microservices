package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/mcpalek/golang-microservices/configloader"

	mssql "github.com/microsoft/go-mssqldb"
)

var insertData = []string{
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (1,'Alice', 'Smith')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (2,'Bob', 'Johnson')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (3,'Charlie', 'Williams')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (4,'David', 'Jones')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (5,'Eva', 'Brown')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (6,'Frank', 'Davis')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (7,'Grace', 'Miller')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (8,'Hannah', 'Wilson')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (9,'Ian', 'Moore')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (10,'Jack', 'Taylor')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (11,'Kathy', 'Anderson')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (12,'Liam', 'Thomas')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (13,'Mia', 'Jackson')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (14,'Nina', 'White')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (15,'Oscar', 'Harris')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (16,'Paul', 'Martin')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (17,'Quincy', 'Thompson')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (18,'Rita', 'Garcia')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (19,'Steve', 'Martinez')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (20,'Tina', 'Roberts')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (21,'Ursula', 'Clark')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (22,'Victor', 'Rodriguez')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (23,'Wendy', 'Lewis')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (24,'Xander', 'Lee')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (25,'Yara', 'Walker')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (26,'Zane', 'Hall')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (27,'Aaron', 'Allen')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (28,'Beth', 'Young')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (29,'Carlos', 'King')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (30,'Diana', 'Scott')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (31,'Edward', 'Green')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (32,'Fiona', 'Adams')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (33,'George', 'Baker')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (34,'Holly', 'Gonzalez')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (35,'Ian', 'Nelson')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (36,'Julia', 'Carter')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (37,'Kevin', 'Mitchell')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (38,'Lilly', 'Perez')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (39,'Mike', 'Robinson')",
	"USE userDB; INSERT INTO Users (id, FirstName, LastName) VALUES (40,'Nancy', 'Jackson')",
}

var (
	db   *sql.DB
	once sync.Once
)

// Retrieves the number of CPU cores from SQL Server
func getSQLServerCPUCount(db *sql.DB) int {
	var cpuCount int
	query := "SELECT cpu_count FROM sys.dm_os_sys_info"
	err := db.QueryRow(query).Scan(&cpuCount)
	if err != nil {
		log.Printf("Error getting CPU count from SQL Server: %v", err)
		return 2 // Default to 2 workers if the query fails
	}
	return cpuCount
}

// Creates the database and table
func setupDatabase(db *sql.DB) {
	_, err := db.Exec("IF NOT EXISTS (SELECT name FROM sys.databases WHERE name = 'userDB') CREATE DATABASE userDB")
	if err != nil {
		log.Fatal("Error creating database:", err)
	}

	_, err = db.Exec("USE userDB; IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='Users' AND xtype='U') CREATE TABLE Users (id INT PRIMARY KEY, FirstName VARCHAR(50), LastName VARCHAR(50))")
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}

// Insert data using a Fan-Out pattern with error handling
func insertDataConcurrently(db *sql.DB, numWorkers int) {
	start := time.Now()
	queryChan := make(chan string, len(insertData))
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for query := range queryChan {
				// Attempt to execute the query with error handling
				success := false
				for attempts := 0; attempts < 3; attempts++ { // Retry up to 3 times
					_, err := db.Exec(query)
					if err != nil {
						if isDuplicateError(err) {
							// Handle duplicate key error (skip this query or log it)
							fmt.Printf("Worker %d: Duplicate record found for query: %s\n", workerID, query)
							break
						}
						// Log other errors and retry
						log.Printf("Worker %d: Error executing %s, attempt %d: %v\n", workerID, query, attempts+1, err)
					} else {
						// Success, log the insertion
						fmt.Printf("Worker %d: Inserted data successfully\n", workerID)
						success = true
						break
					}
					// Add a small delay between retries
					time.Sleep(time.Second)
				}

				// If not successful after retries, log failure
				if !success {
					log.Printf("Worker %d: Failed to insert after retries for query: %s\n", workerID, query)
				}
			}
		}(i)
	}

	// Send queries to workers
	for _, query := range insertData {
		queryChan <- query
	}
	close(queryChan) // No more queries

	wg.Wait() // Wait for all workers

	elapsed := time.Since(start)
	fmt.Printf("Parallel Inserts Execution Time: %v\n", elapsed)
}

// Check if the error is a duplicate key error
func isDuplicateError(err error) bool {
	if err == nil {
		return false
	}

	// Type assert the error to *mssql.Error and check the number field
	if sqlErr, ok := err.(*mssql.Error); ok {
		if sqlErr.Number == 2627 {
			return true
		}
	}
	return false
}

func main() {

	config, err := configloader.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	// once.Do(func() {
	// 	config, err := configloader.LoadConfig()
	// 	if err != nil {
	// 		log.Fatal("Error loading config:", err)
	// 		})
	connString := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;encrypt=disable",
		config.SQLServer.Host, config.SQLServer.Port, config.SQLServer.User, config.SQLServer.Password)
	//fmt.Println("Loaded config:", config)
	start := time.Now()

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	setupDatabase(db)
	elapsed := time.Since(start)
	fmt.Printf("Execution Time: %v\n", elapsed)
	sqlServerCPU := getSQLServerCPUCount(db)
	fmt.Printf("Using %d workers (equal to SQL Server CPU cores)\n", sqlServerCPU)

	insertDataConcurrently(db, sqlServerCPU)

}
