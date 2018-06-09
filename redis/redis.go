// Package redis is a encapsulation for go-redis package
//
// The key is represented by a encoded ID used on database
// it means that decoding the key you will retrieve the ID for the register on database
package redis

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func init() {
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}

// Get retrieve the URL for given encoded ID
func Get(encode string) (value string, err error) {
	value, err = client.Get(encode).Result()
	if err == redis.Nil {
		return
	} else if err != nil {
		log.Fatal(err)
	}

	return
}

// Set saves the given URL for a encoded ID
func Set(encode, url string) (err error) {
	err = client.Set(encode, url, time.Minute*15).Err()
	if err != nil {
		return
	}

	return
}
