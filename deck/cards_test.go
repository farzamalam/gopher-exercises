package deck

import "fmt"

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Queen, Suit: Club})
	fmt.Println(Card{Rank: King, Suit: Diamond})
	fmt.Println(Card{Suit: Joker})
	// Output:
	// Ace of Hearts
	// Two of Spades
	// Queen of Clubs
	// King of Diamonds
	// Joker
}
