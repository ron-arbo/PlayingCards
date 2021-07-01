package deck

import (
	"fmt"
	"math/rand"
)

type Card struct {
	Value int
	Suit string
}

func (c Card) String() string {
	return fmt.Sprintf("%v of %v \n", c.Value, c.Suit)
}

// ByVal implements the sort.Interface for []Card based on the value field
type ByVal []Card
func (a ByVal) Len() int           { return len(a) }
func (a ByVal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVal) Less(i, j int) bool { return a[i].Value < a[j].Value }

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

// Filter filters cards out based on a function that the user must provide
func Filter(f func(c Card) bool, inputDeck []Card) []Card {
	var returnDeck []Card
	for _, v := range inputDeck {
		if !f(v) {
			returnDeck = append(returnDeck, v)
		}
	}
	return returnDeck
}

//Shuffle will shuffle a current deck and return a new deck
func Shuffle(inputDeck []Card) []Card {
	returnDeck := make([]Card, len(inputDeck))

	// Returns a slice of indexes randomly permuted
	perm := rand.Perm(len(inputDeck))
	for i, v := range perm {
		returnDeck[i] = inputDeck[v]
	}
	return returnDeck
}