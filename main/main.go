package main

import deck "github.com/ron-arbo/PlayingCards/deck"

func main() {
	newDeck := deck.New() 
	deck.PrintDeck(newDeck)
}