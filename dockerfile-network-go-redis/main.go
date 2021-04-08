package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name, err := rdb.Get(context.Background(), "name").Result()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, err.Error())
		}
		fmt.Fprintf(w, fmt.Sprintf("Hello %s", name))
	})

	http.ListenAndServe(":8080", nil)
}
