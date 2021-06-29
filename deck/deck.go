package deck

import "fmt"

type Card struct {
	value int
	suit string
}

func New() []Card {
	// Create card slice
	cardDeck := make([]Card, 52)
	suits := [4]string{"spades", "diamonds", "clubs", "hearts"}
	
	// Lopp through suits
	for j := 0; j < 4; j++ {
		//Create cards for each suit
		for i := 0; i < 13; i++ {
			cardDeck = append(cardDeck, Card{i, suits[j]})
		}
	}

	return cardDeck
}

//printCard prints the attribtutes of a given card
func PrintCard(inputCard Card) string {
	return fmt.Sprintf("Value: %v -- Suit: %v", inputCard.value, inputCard.suit)
}

// printDeck prints the contents of a given card slice (deck)
func PrintDeck(inputDeck []Card) string {
	for _, v := range inputDeck {
		PrintCard(v)
	}
	return "hi"
}