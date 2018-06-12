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

// Set saves the given URL for an encoded ID
func Set(key, url string) (err error) {
	err = client.Set(key, url, 30*time.Minute).Err()
	if err != nil {
		log.Println("Fail to save key: " + err.Error())
	}

	return
}

// Get retrieve the URL for given encoded encoded ID
func Get(key string) (result string, err error) {
	result, err = client.Get(key).Result()

	if err == redis.Nil || err != nil {
		return
	}

	return
}
