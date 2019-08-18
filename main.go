package main

import (
	"fmt"
	"github/Blackjack/blackjack"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
