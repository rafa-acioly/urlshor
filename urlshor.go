package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/rafa-acioly/urlshor/database"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", home)
	router.HandleFunc("/short", shortURL).Methods("POST")
	router.HandleFunc("/short/get/{id}", getURL)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
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
		internalServerError("Error trying to read request body; " + err.Error())
	}

	var short struct {
		URL string `json:"url"`
	}
	err = json.Unmarshal(request, &short)
	if err != nil {
		internalServerError("Error trying to unmarshall " + erro.Error())
	}

	// Check if the request have a valid URL
	if _, err = url.ParseRequestURI(short.URL); err != nil {
		log.Println("Invalid URL;" + err.Error())
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}

	// Get the last inserted ID and sum +1 to find out which is the next ID to be inserted on database
	id, err := database.GetLastInsertedID()

	// Generate a encode with base36 on the (last inserted ID + 1)
	encoded := encode36(id)

	// Save the URL and the encode on database
	err = database.Insert(encode, short.URL)
	if err != nil {
		internalServerError("Could not insert register on database." + err.Error())
	}

	// Save the URL and the encoded on redis
	_ := redis.Set(encode, short.URL)

	// return the new encoded URL
	json.NewEncoder(w).Encode(map[string]string{"url": encoded})
}

func internalServerError(w http.ResponseWriter, msg ...string) {
	log.Println(msg)
	http.Error(w, "Internal server error.", http.StatusInternalServerError)
}

func getURL(w http.ResponseWriter, r *http.Request) {
	// Check if the id is on redis and redirect to the URL if found

	// Check if the id is on database and redirect to the URL if found

	// If we do not find the ID, show a 404 page
}
