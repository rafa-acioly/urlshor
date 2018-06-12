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
	port     = 5432
	user     = "postgres"
	password = ""
	dbName   = "urlshor"
)

var (
	dsn = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)

	database *sql.DB
)

func init() {
	database, _ = sql.Open("postgres", "user=postgres dbname=urlshor sslmode=disable")
}

// NextID return the next value for column "id"
func NextID() (uint64, error) {
	var id uint64
	err := database.QueryRow("SELECT nextval('seq_urls_id')").Scan(&id)

	return id, err
}

// Create makes a insert query on database
func Create(id uint64, encode, url string) error {
	query := fmt.Sprintf("INSERT INTO urls VALUES(%d, '%s', '%s', 0)", id, url, encode)
	fmt.Println(query)
	_, err := database.Query(query)

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
