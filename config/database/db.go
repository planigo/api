package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// DB is the underlying database connection
var DB *sql.DB

// Connect initiate the database connection and migrate all the tables
func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := mysql.Config{
		User:                 os.Getenv("MARIADB_ROOT_USERNAME"),
		Passwd:               os.Getenv("MARIADB_ROOT_PASSWORD"),
		DBName:               os.Getenv("MARIADB_DATABASE"),
		Addr:                 os.Getenv("MARIADB_HOST") + ":" + os.Getenv("MARIADB_PORT"),
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	if err = db.Ping(); err != nil {
		fmt.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	DB = db

	fmt.Println("[DATABASE]::CONNECTED")
}
