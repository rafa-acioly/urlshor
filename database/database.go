// Package database is a encapsulation for postgres driver
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	// Package included to connect the database,
	// there is no need to import it on main
	// because the connection is abstracted in this package
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
	stmt, _ := database.Prepare("INSERT INTO urls VALUES($1, $2, $3)")
	_, err := stmt.Exec(id, url, encode)

	return err
}

// Find return an URL for the given encoded key
func Find(encode string) string {
	query := fmt.Sprintf("SELECT url FROM urls WHERE encoded = '%s'", encode)

	var url string
	rows, err := database.Query(query)
	if err != nil {
		log.Fatal("Could not select data from database " + err.Error())
	}

	for rows.Next() {
		rows.Scan(&url)
	}

	return url
}

// IncrementClickCounter add +1 to "clicks" column on given key (encode column)
func IncrementClickCounter(key string) error {
	query := fmt.Sprintf("UPDATE urls SET clicks = clicks + 1 WHERE encoded = '%s'", key)
	_, err := database.Query(query)
	if err != nil {
		return errors.New("Could not increment clicks counter to encode: " + key + err.Error())
	}

	return nil
}
