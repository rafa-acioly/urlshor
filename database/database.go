// Package database is a encapsulation for postgres driver
package database

// GetLastInsertedID return the next value for column "id" using "next_val"
func GetLastInsertedID() (uint64, error) {
	return 123, nil
}

// Insert makes a insert query on database
func Insert(encode, url string) error {
	return nil
}

// Query makes a select statement on database
func Query(encode string) {

}
