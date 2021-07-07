package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title string
	Redis map[string]redis
}

type redis struct {
	Host string
	Port int
}

func main() {
	// config, err := toml.LoadFile(".config")
	// if err != nil {
	// 	fmt.Println("error to load file: ", err)
	// }
	var config Config
	f, err := os.OpenFile(".config", os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, _ := ioutil.ReadAll(f)
	if err := toml.Unmarshal(b, &config); err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("title: ", config.Title)

	config.list()
	config.add()
	config.delete()
	config.update()
	writeFile(f)
	err = toml.NewEncoder(f).Encode(config)
	fmt.Println("redis: ", config.Redis)
	if err != nil {
		panic(err)
	}

	// fmt.Println("master: ", config.Redis["master"].Host)

	// r := redis{Host: "127.0.0.2"}
	// config.Redis["localhost"] = r

	// fmt.Println("slave: ", config.Redis["slave-0"])
	// delete(config.Redis, "slave-4")
	// fmt.Println("slave: ", config.Redis["slave-0"])
	// // r4 := redis{Host: "127.0.0.4"}
	// // config.Redis["slave-4"] = r4

}
func writeFile(f *os.File) {
	f.Truncate(0)
	f.Seek(0, 0)
}

func (c *Config) update() {
	r := c.Redis["host-1"]
	r.Port = 4321
	c.Redis["host-1"] = r
}

func (c *Config) delete() {
	delete(c.Redis, "slave-4")
}

func (c *Config) list() {
	for k, r := range c.Redis {
		fmt.Printf("%s: URL: %s:%d\n", k, r.Host, r.Port)
	}
}

func (c *Config) add() {
	r := redis{
		Host: "192.168.1.1",
		Port: 1234,
	}
	c.Redis["host-1"] = r
}

// func list(config *toml.Tree) {
// 	title := config.Get("Title").(string)
// 	fmt.Println("title: ", title)
// 	redis := config.Get("Redis").(*toml.Tree)
// 	r := redis.Delete("localhost")
// 	fmt.Println("localhost: ", r)
// }
