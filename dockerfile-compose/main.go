package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	redis "github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

func main() {
	// parse env
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	dbhost := os.Getenv("POSTGRES_HOST")
	dbport := os.Getenv("POSTGRES_PORT")
	dbuser := os.Getenv("POSTGRES_USER")
	dbpassword := os.Getenv("POSTGRES_PASS")
	dbname := os.Getenv("POSTGRES_DB_NAME")

	// redis init
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// init postgres
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", dbhost, dbport, dbuser, dbpassword, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// get name
		name, _ := rdb.Get(context.Background(), "name").Result()

		// db ping
		errDB := db.Ping()
		log.Println(errDB)

		json, _ := json.Marshal(map[string]interface{}{
			"redis": name,
			"db":    (errDB == nil),
		})

		fmt.Fprintf(w, string(json))
		return
	})

	http.ListenAndServe(":8080", nil)
}
