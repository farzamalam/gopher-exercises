package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	//Make a csv flag to give manual csv file, then parse it.
	csvFlagName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFlagName)
	if err != nil {
		log.Fatalf("Error while opening the file : %s\n", *csvFlagName)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error while reading file :%s", err)
	}
	problems := parseFile(lines)
	fmt.Println(problems)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s =", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You have scored %d out of %d", correct, len(problems))
}

type problem struct {
	q string
	a string
}

func parseFile(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}
