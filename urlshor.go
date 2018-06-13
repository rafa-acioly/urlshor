package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/rafa-acioly/urlshor/database"
	"github.com/rafa-acioly/urlshor/redis"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", home)
	router.HandleFunc("/{id}", getURL).Methods("GET")
	router.HandleFunc("/short", shortURL).Methods("POST")
	router.HandleFunc("/info/{key}", infoURL).Methods("GET")

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func home(w http.ResponseWriter, r *http.Request) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Printf("No caller information; %v", ok)
	}
	tmpl := template.Must(template.ParseFiles(path.Dir(filename) + "/static/index.html"))
	tmpl.Execute(w, nil)
}

func shortURL(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		internalServerError(w, "Error trying to read request body; "+err.Error())
	}

	var short struct {
		URL string `json:"url"`
	}
	err = json.Unmarshal(request, &short)
	if err != nil {
		internalServerError(w, "Error trying to unmarshall "+err.Error())
	}

	// Check if the request have a valid URL
	if _, err = url.ParseRequestURI(short.URL); err != nil {
		log.Println("Invalid URL;" + err.Error())
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}

	// Get the next id from database to be inserted
	id, err := database.NextID()
	if err != nil {
		log.Fatal("Could not get last inserted ID " + err.Error())
	}

	// Generate a encode with base36
	encoded := encode36(id)

	err = database.Create(id, encoded, short.URL)
	if err != nil {
		internalServerError(w, "Could not insert register on database."+err.Error())
	}

	err = redis.Set(encoded, short.URL)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(map[string]string{"url": encoded})
}

func getURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	value, _ := redis.Get(params["id"])
	// If the encode was not found on redis, search in database
	if len(value) == 0 {
		log.Printf("Cache not found for key: %s", params["id"])

		value = database.Find(decode36(params["id"]))
		// If the value was found on database, bound it to redis
		if len(value) > 0 {
			redis.Set(params["id"], value)
			log.Println("Found on database, key included on cache:", params["id"])
		} else {
			log.Printf("Key not found on database as well: %s", params["id"])
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
	}

	err := database.IncrementClickCounter(params["id"])
	if err != nil {
		log.Println(err.Error())
	}

	http.Redirect(w, r, value, http.StatusFound)
}

func infoURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	url := database.Get(params["key"])

	w.Header().Set("Content-Type", "Application/json")

	if len(url.Encoded) == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(url)
}

func internalServerError(w http.ResponseWriter, msg ...string) {
	log.Println(msg)
	http.Error(w, "Internal server error.", http.StatusInternalServerError)
}
