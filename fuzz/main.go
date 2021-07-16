package main

import (
	"fmt"

	"github.com/sahilm/fuzzy"
)

func main() {
	pattern := "mnr"
	data := []string{"game.cpp", "moduleNameResolver.ts", "my name is_Ramsey"}
	matches := fuzzy.Find(pattern, data)
	fmt.Println("matches:", matches)
	for _, match := range matches {
		for i := 0; i < len(match.Str); i++ {
			if contains(i, match.MatchedIndexes) {
				fmt.Print(fmt.Sprintf(string(match.Str[i])))
			} else {
				fmt.Print(string(match.Str[i]))
			}
		}
		fmt.Println()
	}
}

func contains(needle int, haystack []int) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}
