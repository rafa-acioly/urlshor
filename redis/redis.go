// Package redis is a encapsulation for go-redis package
//
// The key is represented by a encoded ID used on database
// it means that decoding the key you will retrieve the ID for the register on database
package redis

import (
	"log"

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
//
// The HMGET is used to retrieve the "url" key from a hash
func Get(encode string) (result string, err error) {
	val := client.HMGet(encode, "url")

	err = val.Err()
	result = val.Val()[0].(string)

	if err == redis.Nil {
		return
	} else if err != nil {
		log.Fatal(err)
	}

	err = IncrementClickCounter(encode)
	if err != nil {
		return
	}

	return
}

// Set saves the given URL for a encoded ID
//
// the HMSET is used to save a hash on cache so we
// can count how many visits occurs within "clicks" counter
func Set(encode, url string) (err error) {
	hash := map[string]interface{}{
		"url":    url,
		"clicks": 0,
	}
	err = client.HMSet(encode, hash).Err()
	if err != nil {
		return
	}

	return
}

// IncrementClickCounter add +1 to the "click" on given hash
func IncrementClickCounter(encode string) (err error) {
	err = client.HIncrBy(encode, "clicks", 1).Err()
	if err != nil {
		return
	}

	return nil
}
