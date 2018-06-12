// Package database is a encapsulation for postgres driver
package database

import (
	"database/sql"
	"errors"
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
		"user=%s dbname=%s sslmode=disable",
		user, dbName,
	)

	database *sql.DB
)

func init() {
	database, _ = sql.Open("postgres", dsn)
}

// NextID return the next value for column "id"
func NextID() (uint64, error) {
	var id uint64
	err := database.QueryRow("SELECT nextval('seq_urls_id')").Scan(&id)

	return id, err
}

// Create makes a insert query on database
func Create(id uint64, encode, url string) error {
	query := fmt.Sprintf("INSERT INTO urls VALUES(%d, '%s', '%s')", id, url, encode)
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

// IncrementClickCounter add +1 to "clicks" column on given key (encode column)
func IncrementClickCounter(key string) error {
	query := fmt.Sprintf("UPDATE urls SET clicks = clicks + 1 WHERE encode = %s", key)
	_, err := database.Query(query)
	if err != nil {
		return errors.New("Could not increment clicks counter to encode: " + key)
	}

	return nil
}
