package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	pghost := os.Getenv("PG_HOST")
	pgport := os.Getenv("PG_PORT")
	pguser := os.Getenv("PG_USER")
	pgpassword := os.Getenv("PG_PASSWORD")
	pgdbname := os.Getenv("PG_DBNAME")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", pghost, pgport, pguser, pgpassword, pgdbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Successfully connected!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name, err := rdb.Get(context.Background(), "name").Result()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, err.Error())
		}

		err = db.Ping()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, err.Error())
		}

		fmt.Fprintf(w, fmt.Sprintf("Hello %s", name))
	})

	http.ListenAndServe(":8080", nil)
}
