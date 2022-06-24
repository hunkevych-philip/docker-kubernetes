package main

import (
	"context"
	"fmt"
	"hunkevych-philip/docker-kubernetes/datastore/redis"
	"log"
	"net/http"
	"os"
)

var Logger = log.New(os.Stdout, "-> ", 0)

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/crash-test", crash)

	Logger.Println("Starting server on a port 8080")
	Logger.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rdb, err := redis.GetRedisWrapper(ctx)
	if err != nil {
		Logger.Printf("Failed to establish database connection: %s\n", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}

	v, err := rdb.NumberOfVisists(ctx)
	if err != nil {
		Logger.Printf("Database returned an error: %s\n", err.Error())
		w.Write([]byte("Something went wrong"))

		return
	}

	w.Write([]byte(fmt.Sprintf("Number of visits: %d", v)))
}

func crash(w http.ResponseWriter, r *http.Request) {
	Logger.Println("Let's exit!")
	os.Exit(0)
}
