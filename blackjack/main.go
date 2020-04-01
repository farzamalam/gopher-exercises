package main

import (
	"fmt"
	"strings"

	"github.com/farzamalam/gopher-exercises/deck"
)

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	for i := 0; i < 10; i++ {
		card, cards = draw(cards)
		fmt.Println(card)
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	card, cards := cards[0], cards[1:]
	return card, cards
}

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}
