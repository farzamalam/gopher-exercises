package main

import (
	"fmt"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

func main() {
	slice := []struct {
		id   string
		name string
	}{
		{"id1", "foo"},
		{"id2", "bar"},
		{"id3", "baz"},
	}
	idx, _ := fuzzyfinder.Find(slice, func(i int) string {
		return fmt.Sprintf("[%s] %s", slice[i].id, slice[i].name)
	})
	fmt.Println(slice[idx]) // The selected item.
}
