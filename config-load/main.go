package main

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title    string
	App      app
	DB       mysql `toml:"mysql"`
	Redis    map[string]redis
	Releases releases
	Company  Company
}

type app struct {
	Auther  string
	Org     string `toml:"organization"`
	Mark    string
	Release time.Time
}

type mysql struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type redis struct {
	Host string
	Port int
}

type releases struct {
	Release []string
	Tags    [][]interface{}
}

type Company struct {
	Name   string
	Detail detail
}

type detail struct {
	Type string
	Addr string
	ICP  string
}

func main() {
	var config Config
	if _, err := toml.DecodeFile(".config", &config); err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("config: ", config)
}
