package main

import (
	"fmt"
	"github/Blackjack/deck"
	"strings"
)

//Hand represents a hand of deck
type Hand []deck.Card

//String function returns hand of cards seperated wth a ","
func (h Hand) String() string {
	str := make([]string, len(h))
	for i := range h {
		str[i] = h[i].String()
	}
	return strings.Join(str, ",")
}

//DealerString shows the cards of dealer
func (h Hand) DealerString() string {
	return h[0].String() + "**HIDDEN**"
}

//Score function returns the final score
func (h Hand) Score() int {
	minscore := h.minscore()
	if minscore > 11 {
		return minscore
	}
	for _, hand := range h {
		if int(hand.Rank) == 1 {
			return minscore + 10
		}
	}
	return minscore
}

func draw(c []deck.Card) (deck.Card, []deck.Card) {
	return c[0], c[1:]
}
func (h Hand) minscore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//StartGame returns initial state of the game
func StartGame() Gamestate {
	var ret Gamestate
	ret.Deck = deck.New(deck.Deck(3), deck.Shuffle())
	return ret
}

//Deal function deals the cards
func Deal(gs Gamestate) Gamestate {
	var ret Gamestate
	var card deck.Card
	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)
	for i := 1; i <= 2; i++ {
		for _, hand := range []*Hand{&ret.Player, &ret.Dealer} {
			card, ret.Deck = draw(ret.Deck)
			*hand = append(*hand, card)
		}

	}
	ret.Gstate = playersTurn
	return ret
}

//Clone function copy's a Gamestate data into another
func Clone(gs Gamestate) Gamestate {
	ret := Gamestate{
		Deck:   make([]deck.Card, len(gs.Deck)),
		Gstate: gs.Gstate,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}

//Hit operates when a function opts to hit
func Hit(gs Gamestate) Gamestate {
	ret := Clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret

}

//Stand operates when a function opts to stand
func Stand(gs Gamestate) Gamestate {
	ret := Clone(gs)
	ret.Gstate++
	return ret
}

func main() {
	gs := StartGame()

	var input string
	for input != "s" {
		fmt.Println("Player:", gs.Player.String())
		fmt.Println("Dealer:", gs.Dealer.DealerString())
		fmt.Println("what will you do now, (h)it,(s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			Hit(gs)
		case "s":
			Stand(gs)
		default:
			fmt.Println("wrong choice")
		}

	}

	for gs.Gstate == dealerTurn {

		if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.minscore() == 7) {
			Hit(gs)
		} else {
			Stand(gs)
		}
	}
	EndGame(gs)
}

//EndGame shows the final Score
func EndGame(gs Gamestate) {
	dscore := gs.Dealer.Score()
	pscore := gs.Player.Score()

	fmt.Println("===FINAL HAND===")
	fmt.Println("Player:", gs.Player.String(), "\nscore:", pscore)
	fmt.Println("Dealer:", gs.Dealer.String(), "\nscore:", dscore)
	switch {
	case dscore > 21:
		fmt.Println("Dealer Busted")
	case pscore > 21:
		fmt.Println("You Busted")
	case pscore < dscore:
		fmt.Println("Player: You lose")
	case dscore < pscore:
		fmt.Println("Player:You Win")
	case dscore == pscore:
		fmt.Println("Draw")

	}
}

type gamestate uint8

const (
	playersTurn gamestate = iota
	dealerTurn
	endTurn
)

//Gamestate is a struct which has all the features of a game at a condition
type Gamestate struct {
	Deck   []deck.Card
	Gstate gamestate
	Player Hand
	Dealer Hand
}

//CurrentPlayer function returns the current player in action
func (gs Gamestate) CurrentPlayer() *Hand {
	if gs.Gstate == playersTurn {
		return &gs.Player
	}
	return &gs.Dealer

}
