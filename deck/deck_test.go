package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	myCard := Card{4, "spades"}
	fmt.Println(myCard)
	// Output: 4 of spades
}

func TestShuffle(t *testing.T) {
	deck1 := New()
	deck2 := Shuffle(deck1)

	// Assuming it's impossible to "shuffle" a deck and 
	// get the same exact result...
	if len(deck1) != len(deck2) {
		t.Fatal("Decks are not the same length")
	}
	for i, _ := range deck1 {
		if deck1[i] != deck2[i] {
			return 
		}	
	}
	t.Fatal("Decks were the same")
}

func TestFilter(t *testing.T) {
	deck1 := New()
	f := func(c Card) bool{
		return c.Value == 2
	}
	deck1 = Filter(f, deck1)
	for _, v := range deck1 {
		if v.Value == 2 {
			t.Fatal("Did not filter out the 2s")
		}
	}
}