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

// URLInfo represents all fields on url table
type URLInfo struct {
	ID        int    `json:"-"`
	URL       string `json:"url"`
	Encoded   string `json:"encoded"`
	Clicks    int    `json:"clicks"`
	CreatedAt string `json:"created_at"`
}

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
func Create(id uint64, encoded, url string) error {
	stmt, _ := database.Prepare("INSERT INTO urls VALUES($1, $2, $3)")
	_, err := stmt.Exec(id, url, encoded)

	return err
}

// Find return an URL for the given encoded key
func Find(id uint64) string {
	query := fmt.Sprintf("SELECT url FROM urls WHERE id = '%d'", id)

	var url string
	rows, err := database.Query(query)
	if err != nil {
		log.Fatal("Could not select data from database " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&url)
	}

	return url
}

// Get return a complete data register from database for the given encoded key
func Get(key string) URLInfo {
	query := fmt.Sprintf("SELECT * FROM urls WHERE encoded = '%s'", key)

	rows, err := database.Query(query)

	if err != nil {
		log.Fatal("Could not select data from database: " + err.Error())
	}
	defer rows.Close()

	var url URLInfo
	for rows.Next() {
		rows.Scan(&url.ID, &url.URL, &url.Encoded, &url.Clicks, &url.CreatedAt)
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
