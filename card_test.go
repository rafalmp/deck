package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Spade})
	fmt.Println(Card{Rank: Six, Suit: Heart})
	fmt.Println(Card{Rank: Jack, Suit: Diamond})
	fmt.Println(Card{Rank: King, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Spades
	// Six of Hearts
	// Jack of Diamonds
	// King of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in the new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	want := Card{Suit: Club, Rank: Ace}
	if cards[0] != want {
		t.Errorf("Expected %v as the first card, got %v", want, cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	want := Card{Suit: Club, Rank: Ace}
	if cards[0] != want {
		t.Errorf("Expected %v as the first card, got %v", want, cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(2))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 2 {
		t.Error("Expected 2 jokers, actual:", count)
	}
}
