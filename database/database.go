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

	// createTable := `CREATE TABLE IF NOT EXISTS orders (
	// id SERIAL PRIMARY KEY,
	// order_id VARCHAR(256) UNIQUE NOT NULL,
	// customer_id VARCHAR(256) UNIQUE NOT  NULL,
	// product_id TEXT[] UNIQUE NOT NULL,
	// Quantity INTEGER,
	// hotel_id VARCHAR(256) UNIQUE NOT NULL,
	// order_address VARCHAR(256),
	// order_location VARCHAR(256),
	// Order_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	// order_status VARCHAR(256),
	// created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	// updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	// )`

	// _, err = repository.DB.Exec(createTable)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Table created successfully")
}
