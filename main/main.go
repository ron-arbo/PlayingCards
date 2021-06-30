package main

import "github.com/ron-arbo/PlayingCards/deck"

func main() {
	newDeck := deck.New()
	deck.PrintDeck(newDeck)
	
	newDeck = deck.FilterCards(newDeck, 4)
	deck.PrintDeck(newDeck)
}