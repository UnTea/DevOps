package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "",
	DB:       0,
})

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Time).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Time(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	value, result := client.Get(ctx, "time").Result()
	if result == redis.Nil {
		value = time.Now().Add(3 * time.Hour).Format("02-01-2006 15:04:05")
		err := client.Set(ctx, "time", value, 10*time.Second).Err()
		if err != nil {
			fmt.Println(err)
		}
	}

	_, err := fmt.Fprintf(w, "%v", value)
	if err != nil {
		log.Fatalf("Error occured while handle response %v", err)
	}

	log.Println(value)
}
