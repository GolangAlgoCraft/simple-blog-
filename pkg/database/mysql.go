// mysql.go
package database

import (
	"database/sql" // SQL package for database interaction
	"log"          // Logging package

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

var DB *sql.DB // Declare a global variable for the database connection

// InitMySQL initializes the MySQL connection
func InitMySQL(dataSourceName string) error {
	var err error
	// Open a connection to the MySQL database
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	// Ping the database to ensure the connection is valid
	return DB.Ping()
}

// Close closes the database connection
func Close() {
	if err := DB.Close(); err != nil {
		log.Println("Error closing database connection:", err)
	}
}
