package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: ace, Suit: spade})

	//output:
	//ace of spades
}
func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("wrong number of cards in new deck")
	}
}
func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	ex := Card{Suit: spade, Rank: ace}
	if cards[0] != ex {
		t.Error("First card should be ace of spades")
	}
}
func TestShuffle(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == joker {
			count++
		}
	}
	if count != 3 {
		t.Error("expected 3 received:", count)
	}
}
func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == two || card.Rank == three
	}
	Cards := New(Filter(filter))
	for _, c := range Cards {
		if c.Rank == two || c.Rank == three {
			t.Error("Deck is not filtered properly")
		}
	}
}
func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != (4 * 13 * 3) {
		t.Error("wrong number of crds in deck")
	}
}
