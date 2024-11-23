package config

import (
	"log"

	"github.com/jmoiron/sqlx" // Import the sqlx library
	_ "github.com/lib/pq"     // PostgreSQL driver
)

var DB *sqlx.DB

// Connect initializes the database connection
func Connect() *sqlx.DB {
	// Build the connection string
	psqlInfo := "user=postgres.fhsvbficbllrmkfpqmut password=KnowV@tion24 host=aws-0-ap-south-1.pooler.supabase.com port=6543 dbname=postgres"

	// Connect to the database
	var err error
	DB, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Successfully connected to the database!")
	return DB
}

// Close disconnects the database connection
func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
		log.Println("Database connection closed.")
	}
}
