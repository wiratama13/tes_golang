package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"gith"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

// Employee struct untuk menyimpan data employee
type Employee struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	TerminationDate sql.NullString `json:"termination_date"`
}

func main() {
	// Konfigurasi koneksi database
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query database
	rows, err := db.Query("SELECT * FROM employee WHERE TerminationDate IS NULL ORDER BY LastName, FirstName ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Slice untuk menyimpan hasil query
	var employees []Employee

	// Iterasi hasil query
	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Email, &emp.TerminationDate)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, emp)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Mengubah slice menjadi JSON
	jsonData, err := json.Marshal(employees)
	if err != nil {
		log.Fatal(err)
	}

	// Mencetak hasil JSON
	fmt.Println(string(jsonData))
}
