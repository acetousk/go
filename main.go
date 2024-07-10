package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Connection string
	connStr := "host=localhost port=5432 user=moqui password=moqui dbname=moqui sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Query the user_account table
	rows, err := db.Query("SELECT email_address FROM user_account")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through the result set
	var emailAddresses []string
	for rows.Next() {
		var email sql.NullString
		if err := rows.Scan(&email); err != nil {
			log.Fatal(err)
		}
		if email.Valid {
			emailAddresses = append(emailAddresses, email.String)
		} else {
			emailAddresses = append(emailAddresses, "NULL")
		}
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Print the results
	fmt.Println("Email addresses:")
	for _, email := range emailAddresses {
		fmt.Println(email)
	}
}