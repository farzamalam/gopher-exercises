package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type file struct {
	name string
	path string
}

func main() {

	dir := "sample"

	var toRename []file

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if _, err := match(info.Name()); err == nil && !info.IsDir() {
			toRename = append(toRename, file{
				name: info.Name(),
				path: path,
			})
		}
		return nil
	})
	for _, f := range toRename {
		fmt.Println(f)
	}

	for _, orig := range toRename {
		var n file
		var err error
		n.name, err = match(orig.name)
		if err != nil {
			fmt.Println("Error while renaming ", n.name)
		}
		n.path = filepath.Join(dir, n.name)
		err = os.Rename(orig.path, n.path)
		if err != nil {
			panic(err)
		}
	}

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
		return "", fmt.Errorf("%s did not match our pattern", fileName)
	}
	return fmt.Sprintf("%s - %d.%s", strings.Title(name), number, ext), nil

}
