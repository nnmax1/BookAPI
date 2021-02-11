package main 

import (
"fmt"
"github.com/gorilla/mux"
"github.com/nnmax1/BookAPI/datastore"
"log"
"net/http"
"time"
"os"
"encoding/json"
)

func getAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
		data := books
		if data != nil {
			b, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			return
		}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}


func GetPort() string {
 	var port = os.Getenv("PORT")
 	// Set a default port if there is nothing in the environment
 	if port == "" {
 		port = "8080"
 		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
 	}
 	return ":" + port
}


var (
	books datastore.BookStore
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	books = &datastore.Books{}
    books.Initialize()
}


func main() {

	r := mux.NewRouter()
	log.Println("Book Api")
	api := r.PathPrefix("/v1").Subrouter()
	//sub router 1
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "v1")
	})

	//GET Request all data /v1/books
	api.HandleFunc("/bookData", getAll).Methods(http.MethodGet)


	fmt.Println("listening...")
	 	err := http.ListenAndServe(GetPort(), r)
	 	if err != nil {
	 		log.Fatal("ListenAndServe: ", err)
	}
}