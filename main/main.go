package main

import (
	"fmt"

	"github.com/ron-arbo/PlayingCards/deck"
)

func main() {
	newDeck := deck.New()
	// deck.PrintDeck(newDeck)
	
	newDeck = deck.Filter(newDeck, 4)
	// deck.PrintDeck(newDeck)

	fmt.Println(newDeck)
}