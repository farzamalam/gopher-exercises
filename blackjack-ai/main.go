package main

import (
	"fmt"

	"github.com/farzamalam/gopher-exercises/blackjack-ai/blackjack"
)

func main() {
	opts := blackjack.Options{
		Decks:           2,
		Hands:           1,
		BlackjackPayout: 1.5,
	}
	game := blackjack.New(opts)
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
