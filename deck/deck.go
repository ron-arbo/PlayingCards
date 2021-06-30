package deck

import "fmt"

type Card struct {
	Value int
	Suit string
}

func New() []Card {
	// Create card slice
	cardDeck := make([]Card, 52)
	suits := [4]string{"spades", "diamonds", "clubs", "hearts"}
	
	// Loop through suits
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
	return fmt.Sprintf("Value: %v -- Suit: %v", inputCard.Value, inputCard.Suit)
}

// printDeck prints the contents of a given card slice (deck)
func PrintDeck(inputDeck []Card) {
	for _, v := range inputDeck {
		fmt.Println(PrintCard(v))
	}
}

// filterCards filters out cards that don't meet a given value threshold
func FilterCards(inputDeck []Card, thresh int) []Card {
	newDeck := make([]Card, 0)
	for _, v := range inputDeck {
		if v.Value >= thresh {
			newDeck = append(newDeck, v)
		}	
	}
	return newDeck
}