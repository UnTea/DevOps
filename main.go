package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", printTime).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func printTime(w http.ResponseWriter, r *http.Request) {
	h, m, s := time.Now().Clock()

	_, err := fmt.Fprintf(w, "%v:%v:%v", h, m, s)
	if err != nil {
		log.Fatalf("Error occured while handle response %v", err)
	}
}
