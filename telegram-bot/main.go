package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
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
		return
	}
	if err := sayHello(body.Message.Chat.ID); err != nil {
		log.Println("Error in sending reply : ", err)
		return
	}
	log.Println("Reply sent.")

}

type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func sayHello(chatID int64) error {
	fmt.Println("ChatID : ", chatID)

	// Create a request body.
	req := &sendMessageReqBody{
		ChatID: chatID,
		Text:   "Hello, How are you ?",
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	_ = godotenv.Load()
	// Send  a post request with your token.
	telegramToken := os.Getenv("telegram_token")
	url := "https://api.telegram.org/bot" + telegramToken + "/sendMessage"
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Println("Error in Post : ", err)
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("Unexpected Status " + res.Status)
	}
	return nil
}
func main() {
	log.Println("Starting server at 3000")
	log.Fatal(http.ListenAndServe(":3000", http.HandlerFunc(Handler)))
}
