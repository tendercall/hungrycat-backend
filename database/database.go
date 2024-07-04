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

	// createTable := `CREATE TABLE IF NOT EXISTS restaurant (
	// id SERIAL PRIMARY KEY,
	// hotel_id VARCHAR(256) UNIQUE NOT NULL,
	// hotel_name VARCHAR(256),
	// description VARCHAR(256),
	// address VARCHAR(256),
	// location VARCHAR(256),
	// phone_number VARCHAR(256),
	// email VARCHAR(256),
	// website VARCHAR(256),
	// menu TEXT[],
	// profile_image VARCHAR(256),
	// open_time VARCHAR(256),
	// close_time VARCHAR(256),
	// ratings INTEGER,
	// created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	// updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	// )`

	// _, err = repository.DB.Exec(createTable)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Table created successfully")
}
