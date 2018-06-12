// Package redis is a encapsulation for go-redis package
//
// The key is represented by a encoded ID used on database
// it means that decoding the key you will retrieve the ID for the register on database
package redis

import (
	"errors"
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

// Set saves the given URL for a encoded ID
//
// the HMSET is used to save a hash on cache so we
// can count how many visits occurs with "clicks" counter
// the default value to "clicks" counter is zero.
func Set(key, url string) (err error) {
	hash := map[string]interface{}{
		"url":    url,
		"clicks": 0,
	}

	err = client.HMSet(key, hash).Err()
	if err != nil {
		log.Println("Fail to save hash: " + err.Error())
	}

	err = client.Expire(key, time.Minute*10).Err()
	if err != nil {
		log.Println("Fail to set expire time to hash: " + err.Error())
	}

	return
}

// Get retrieve the URL for given encoded ID
//
// The HMGET is used to retrieve the "url" key from a hash
func Get(key string) (result string, err error) {
	val := client.HMGet(key, "url")

	err = val.Err()

	if err == redis.Nil || err != nil {
		return "", errors.New("Key not found")
	}

	// TODO: Apply validation to key not found
	// there is no error above so this key is empty and the converstion crash
	result = val.Val()[0].(string)

	err = IncrementClickCounter(key)
	if err != nil {
		log.Printf("Could not increment the click counter to key: %s - %s", key, err.Error())
	}

	return
}

// IncrementClickCounter add +1 to the "click" on given hash key
func IncrementClickCounter(key string) (err error) {
	err = client.HIncrBy(key, "clicks", 1).Err()
	if err != nil {
		return
	}

	return nil
}
