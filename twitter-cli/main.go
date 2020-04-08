package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var keys struct {
		Key    string `json:"consumer_key"`
		Secret string `json:"consumer_secret"`
	}
	f, err := os.Open("keys.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	d := json.NewDecoder(f)
	d.Decode(&keys)
	fmt.Printf("%+v\n", keys)
}
