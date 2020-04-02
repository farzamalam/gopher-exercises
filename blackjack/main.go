package main

import (
	"fmt"
	"strings"

	"github.com/farzamalam/gopher-exercises/deck"
)

// main func is used to orcastrate all methods and funcs.
// it initializes and uses same gs across all the funcs.
func main() {
	var gs GameState
	gs = Shuffle(gs)

	for i := 0; i < 10; i++ {
		gs = Deal(gs)

		var input string
		for gs.State == StatePlayerTurn {
			fmt.Println("Player:", gs.Player)
			fmt.Println("Dealer:", gs.Dealer.dealerString())
			fmt.Println("What will you do? (h)it, (s)tand")
			fmt.Scanf("%s\n", &input)
			switch input {
			case "h":
				gs = Hit(gs)
			case "s":
				gs = Stand(gs)
			default:
				fmt.Println("Invalid option:", input)
			}
		}

		for gs.State == StateDealerTurn {
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.minScore() != 17) {
				gs = Hit(gs)
			} else {
				gs = Stand(gs)
			}
		}

		gs = EndHand(gs)
	}
}

// Score is used to return the total score, it factors ACE.
func (h Hand) Score() int {
	minScore := h.minScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

// minScore calculates the total score by considering ace as 1
func (h Hand) minScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

// min is used to return min of two values
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// draw returns a card and a deck of cards other than removed card.
func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	card, cards := cards[0], cards[1:]
	return card, cards
}

// Hand type is used to represent a current state of hand.
type Hand []deck.Card

// It prints the hand with comma seperated string
func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

// dealerString is used to show only one card of dealer.
func (h Hand) dealerString() string {
	return h[0].String() + ", ** Hidden **"
}

// State represents who has the turn.
type State uint8

// GameState type represents the current state of a game.
type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

//consts to decide the turns
const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

//clone is used to deepClone of two GameState structs.
func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
		State:  gs.State,
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}

// Current player returns which player's hand
func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StateDealerTurn:
		return &gs.Dealer
	case StatePlayerTurn:
		return &gs.Player
	default:
		panic("No body knows whos turn is this.")
	}
}

// Shuffle creates a new 3 decks and shuffles them.
func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	return ret
}

// Deal is used to deal the card to player and dealer.
func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)

	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)

		ret.State = StatePlayerTurn
	}
	return ret
}

// Hit is used to draw one more card from the deck for dealer or player
func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

// Stand just changes the playerState
func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

// End hand is used to print the score and winner or losser.ls

func EndHand(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", ret.Player, "\nScore:", pScore)
	fmt.Println("Dealer:", ret.Dealer, "\nScore:", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You lose")
	case dScore == pScore:
		fmt.Println("Draw")
	}
	fmt.Println()

	ret.Player = nil
	ret.Dealer = nil
	return ret
}
