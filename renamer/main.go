package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fileName := "birthday_001.txt"
	newName, err := match(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(newName)
}

func match(fileName string) (string, error) {
	// Birthday - 1.txt
	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%s - %d.%s", strings.Title(name), number, ext), nil

}
