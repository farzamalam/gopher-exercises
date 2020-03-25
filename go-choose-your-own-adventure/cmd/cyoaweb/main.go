package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/farzamalam/gopher-exercises/go-choose-your-own-adventure/cyoa"
)

func main() {
	fileName := flag.String("file", "gopher.json", "the JSON file with cyoa story")
	flag.Parse()
	fmt.Println("fileName : ", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("Error while opening file : ", err)
		os.Exit(1)
	}
	story, err := cyoa.JsonDecode(f)
	if err != nil {
		fmt.Println("Error while decoding file : ", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", story)
}
