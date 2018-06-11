// Package database is a encapsulation for postgres driver
package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 8080
	user     = "postgres"
	password = ""
	dbName   = "urlshor"
)

var dsn = fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s",
	host, port, user, password, dbName,
)

var database *sql.DB

func init() {
	database, _ = sql.Open("postgres", dsn)
}

// GetLastInsertedID return the next value for column "id" using "next_val"
func GetLastInsertedID() (id uint64, err error) {
	err = database.QueryRow("SELECT last_value FROM seq_urls").Scan(&id)
	if err != nil {
		log.Fatal("Could not get the last inserted ID " + err.Error())
	}

	return
}

// Create makes a insert query on database
func Create(encode, url string) error {
	query := fmt.Sprintf("INSERT INTO urls(url, clicks) VALUES(%s, %s)", encode, url)
	_, err := database.Query(query)
	if err != nil {
		log.Fatal("Could not insert register in database " + err.Error())
	}

	return err
}

// Find makes a select statement on database
func Find(encode string) string {
	query := fmt.Sprintf("SELECT url FROM urls WHERE encode = %s", encode)
	var url string

	err := database.QueryRow(query).Scan(&url)
	if err != nil {
		log.Fatal("Could not select data from database " + err.Error())
	}

	return url
}
