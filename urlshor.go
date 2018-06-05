package urlshort

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/short", shortURL)
	router.HandleFunc("/short/get/{id}", getURL)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func shortURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Check if the request have a valid URL
	// if not valid return a http.StatusNotAllowed
	if _, err := url.ParseRequestURI(params["url"]); err != nil {
		w.WriteHeader(http.StatusNotAllowed)
	}

	// Get the last inserted ID and sum +1 to find out which is the next ID to be inserted on database

	// Generate a encode with base36 on the (last inserted ID + 1)

	// Save the URL and the encode on database

	// Save the URL and the encoded on redis

	// return the new encoded URL
}

func getURL(w http.ResponseWriter, r *http.Request) {
	// Check if the id is on redis and return the url if found

	// Check if the id is on database and return the url if found

	// If we do not find the ID, show a 404 page

	// redirect to the URL
}
