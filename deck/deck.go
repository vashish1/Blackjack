//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//Suit ....
type Suit uint8

//Rank .....
type Rank uint8

const (
	//Spade ....
	spade Suit = iota
	diamond
	club
	heart
	joker
)

const (
	//_ .......
	_ Rank = iota
	ace
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
)

var suits = [...]Suit{spade, heart, club, diamond}

//Card ....
type Card struct {
	Suit
	Rank
}

const (
	minRank = ace
	maxRank = king
)

func (c Card) String() string {
	if c.Suit == joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

//New is a variadic function
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

//DefaultSort sorts the deck of card
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

//Less is a supportive function
func Less(c []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(c[i]) < absRank(c[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

//Shuffle shuffles the deck of card
func Shuffle() func([]Card) []Card {
	return func(c []Card) []Card {
		ret := make([]Card, len(c))
		r := rand.New(rand.NewSource(time.Now().Unix()))
		perm := r.Perm(len(c))
		for i, j := range perm {
			ret[i] = c[j]
		}
		return ret
	}
}

//Jokers add jokers to the deck
func Jokers(n int) func([]Card) []Card {
	return func(c []Card) []Card {
		for i := 1; i <= n; i++ {
			c = append(c, Card{
				Rank: Rank(i),
				Suit: joker,
			})
		}
		return c
	}
}

//Filter functio filters the deck w.r.t condition
func Filter(f func(Card) bool) func([]Card) []Card {
	return func(c []Card) []Card {
		var ret []Card
		for _, card := range c {
			if !f(card) {
				ret = append(ret, card)

			}
		}
		return ret

	}
}

//Deck returns a deck with multiple deck
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 1; i <= n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
