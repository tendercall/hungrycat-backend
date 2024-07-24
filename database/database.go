package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"hungerycat-backend.com/main/services/repository"
)

func Initdb() {
	// PostgreSQL connection parameters
	const (
		host     = "ep-frosty-moon-a2rk11dy.eu-central-1.pg.koyeb.app"
		port     = 5432
		user     = "koyeb-adm"
		password = "SV1GydCIqM9P"
		dbname   = "koyebdb"
	)

	// Construct the connection string
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

	// Attempt to connect to the database
	var err error
	repository.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = repository.DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Database connection established")

	// createTable := `CREATE TABLE IF NOT EXISTS logs (
	// id SERIAL PRIMARY KEY,
	// function VARCHAR(256),
	// log_message VARCHAR(256),
	// customer_id VARCHAR(256),
	// device_id VARCHAR(256),
	// created_date TIMESTAMP NOT NULL DEFAULT NOW(),
	// updated_date TIMESTAMP NOT NULL DEFAULT NOW()
	// )`

	// _, err = repository.DB.Exec(createTable)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Table created successfully")

	//Example ALTER TABLE statement to add a new column
	// 	query := `
	// 		ALTER TABLE rating
	// 		ALTER COLUMN rating INTEGER;
	// 	`

	// 	// Execute the ALTER TABLE statement
	// 	_, err = repository.DB.Exec(query)
	// 	if err != nil {
	// 		log.Fatalf("Error executing ALTER TABLE statement: %v\n", err)
	// 	}

	// 	fmt.Println("Column added successfully.")
	// }
}
