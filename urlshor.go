package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", home)
	router.HandleFunc("/short", shortURL)
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
		fmt.Println("No caller information")
	}
	tmpl := template.Must(template.ParseFiles(path.Dir(filename) + "/static/index.html"))
	tmpl.Execute(w, nil)
}

func shortURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Check if the request have a valid URL
	if _, err := url.ParseRequestURI(params["url"]); err != nil {
		w.WriteHeader(http.StatusForbidden)
	}

	// Get the last inserted ID and sum +1 to find out which is the next ID to be inserted on database

	// Generate a encode with base36 on the (last inserted ID + 1)

	// Save the URL and the encode on database

	// Save the URL and the encoded on redis

	// return the new encoded URL
}

func getURL(w http.ResponseWriter, r *http.Request) {
	// Check if the id is on redis and redirect to the URL if found

	// Check if the id is on database and redirect to the URL if found

	// If we do not find the ID, show a 404 page
}
