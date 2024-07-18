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

	// createTable := `CREATE TABLE IF NOT EXISTS delivery_boy (
	// id SERIAL PRIMARY KEY,
	// name VARCHAR(256),
	// phone_number VARCHAR(256),
	// db_id VARCHAR(256),
	// Location VARCHAR(256),
	// latitude VARCHAR(256),
	// longitude VARCHAR(256),
	// total_payment INTEGER,
	// total_orders INTEGER,
	// profile_image VARCHAR(256),
	// created_date TIMESTAMP NOT NULL DEFAULT NOW(),
	// updated_date TIMESTAMP NOT NULL DEFAULT NOW()
	// )`

	// _, err = repository.DB.Exec(createTable)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Table created successfully")
}
