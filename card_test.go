package deck

import (
	"fmt"
	"math/rand"
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

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Errorf("%v present but should be filtered out", c)
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	want := 13 * 4 * 3
	got := len(cards)
	if want != got {
		t.Errorf("Expected %d cards, got %d\n", want, got)
	}
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	// first call to shuffleRand.Perm(52) should yield [40 35 ...]
	shuffleRand = rand.New(rand.NewSource(0))

	orig := New()
	first := orig[40]
	second := orig[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected the first card to be %v, got %v\n", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the second card to be %v, got %v\n", second, cards[1])
	}
}
