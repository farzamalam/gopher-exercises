package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	toml "github.com/pelletier/go-toml"
)

type Config struct {
	Title string
	Redis map[string]redis
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

	f, err := os.OpenFile(".config", os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err = toml.Unmarshal(b, &config); err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("config: ", config)
	fmt.Println("title: ", config.Title)
	fmt.Println("master: ", config.Redis["master"].Host)

	// r := redis{Host: "127.0.0.2"}
	// config.Redis["localhost"] = r

	// fmt.Println("slave: ", config.Redis["slave-0"])
	// delete(config.Redis, "slave-4")
	// fmt.Println("slave: ", config.Redis["slave-0"])
	// // r4 := redis{Host: "127.0.0.4"}
	// // config.Redis["slave-4"] = r4
	// // err = toml.NewEncoder(f).Encode(config)
	// fmt.Println("redis: ", config.Redis)
	// if err != nil {
	// 	panic(err)
	// }
}
