package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
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
	//fmt.Printf("%+v\n", keys)

	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(keys.Key, keys.Secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	//io.Copy(os.Stdout, res.Body)

	var token oauth2.Token
	d = json.NewDecoder(res.Body)
	err = d.Decode(&token)
	if err != nil {
		panic(err)
	}
	//fmt.Println("Token ", token)

	ctx := context.Background()
	var conf oauth2.Config
	tclient := conf.Client(ctx, &token)
	usernames, err := retweeters(tclient, "1221142432177221634")
	fmt.Println(usernames)
	fmt.Println("len : ", len(usernames))
}

func retweeters(client *http.Client, tweetID string) ([]string, error) {
	url := fmt.Sprintf("https://api.twitter.com/1.1/statuses/retweets/%s.json", tweetID)
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var retweets []struct {
		User struct {
			ScreenName string `json:"screen_name"`
		} `json:"user"`
	}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&retweets)
	if err != nil {
		return nil, err
	}
	username := make([]string, 0, len(retweets))
	for _, retweet := range retweets {
		username = append(username, retweet.User.ScreenName)
	}
	return username, nil
}
