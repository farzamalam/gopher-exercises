package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Info struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

const token = ""
const url = "https://ipinfo.io"

func main() {
	ip := "103.163.101.35"

	url := fmt.Sprintf("%s/%s?token=%s", url, ip, token)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while calling rest api: ", err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	var info Info

	err = json.Unmarshal([]byte(data), &info)
	if err != nil {
		fmt.Println("Error while unmarshalling: ", err)
	}
	fmt.Printf("%+v\n", info)
}
