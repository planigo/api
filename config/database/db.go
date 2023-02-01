package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Connect() (*sql.DB, error) {

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

	if err = db.Ping(); err != nil {
		fmt.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	fmt.Println("[DATABASE]::CONNECTED")
	return db, nil
}
