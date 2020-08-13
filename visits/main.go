package main

import (
	"golang.org/x/net/context"

	"fmt"
	"log"
	"net/http"
	"strconv"

	redis "github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var client *redis.Client
var visit int
var ctx context.Context

func main() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis-server:6379",
		Password: "",
		DB:       0,
	})
	ctx = context.Background()

	pong, err := client.Ping(ctx).Result()
	visit = 0
	fmt.Println(pong, err)
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	log.Println("Server started at : 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func home(w http.ResponseWriter, r *http.Request) {
	val, err := client.Get(ctx, "visit").Result()
	if err != nil {
		log.Println("Error in home getting value : ", err)
	}
	v, _ := strconv.Atoi(val)
	v++
	fmt.Fprintln(w, "Number of times visited : ", v)
	err = client.Set(ctx, "visit", v, 0).Err()
	if err != nil {
		log.Println("Error in home setting value : ", err)
	}
}
