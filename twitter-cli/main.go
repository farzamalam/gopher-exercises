package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
)

func main() {

	keyFile := flag.String("key", "keys.json", "The file that contains the key and secret")
	tweetID := flag.String("tweet", "1221142432177221634", "The tweet id that is used to get the information.")
	userFile := flag.String("users", "users.csv", "The file that contains the usernames of all the people who retweeted.")
	_ = userFile
	flag.Parse()

	key, secret, err := keys(*keyFile)
	if err != nil {
		panic(err)
	}
	client, err := twitterClient(key, secret)
	if err != nil {
		panic(err)
	}
	usernames, err := retweeters(client, *tweetID)
	fmt.Println(usernames)
	fmt.Println("len : ", len(usernames))
}

func keys(keyFile string) (key, secret string, err error) {
	var keys struct {
		Key    string `json:"consumer_key"`
		Secret string `json:"consumer_secret"`
	}
	f, err := os.Open(keyFile)
	if err != nil {
		return "", "", nil
	}
	defer f.Close()
	d := json.NewDecoder(f)
	d.Decode(&keys)
	return keys.Key, keys.Secret, nil
}

func twitterClient(key, secret string) (*http.Client, error) {
	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(key, secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	//io.Copy(os.Stdout, res.Body)

	var token oauth2.Token
	d := json.NewDecoder(res.Body)
	err = d.Decode(&token)
	if err != nil {
		return nil, err
	}
	//fmt.Println("Token ", token)

	ctx := context.Background()
	var conf oauth2.Config
	return conf.Client(ctx, &token), nil
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
