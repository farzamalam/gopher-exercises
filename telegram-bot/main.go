package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type webHookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	body := &webHookReqBody{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		log.Println("Unable to parse the request body.")
		return
	}
	if !strings.Contains(strings.ToLower(body.Message.Text), "hi") {

	}

}

func main() {

}
