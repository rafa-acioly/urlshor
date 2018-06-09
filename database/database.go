// Package database is a encapsulation for postgres driver
package database

// GetLastInsertedID return the next value for column "id" using "next_val"
func GetLastInsertedID() (uint64, error) {
	return 123, nil
}

// Create makes a insert query on database
func Create(encode, url string) error {
	return nil
}

// Find makes a select statement on database
func Find(encode string) {

}
