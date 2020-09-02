package main

func main() {
	// var card string = "Ace Of Spades"
	// card := "Ace Of Spades"
	// card = "Five Of Diamonds"
	// card := newCard()
	// fmt.Println(card)

	// Slices (like python lists
	// card := []string{newCard(), "Ace of Diamonds"}
	// card := deck{newCard(), "Ace of Diamonds"}
	// card = append(card, "Six of spades")
	// for i, c := range card {
	// 	fmt.Println(i, c)
	// }

	//card.print()

	// card := newDeck()
	// hand, remainingCards := deal(card, 5)
	// hand.print()
	// remainingCards.print()

	//card := newDeck()
	//card.toString()
	// card.saveToFile("myCards")

	// card := newDeckFromFile("myCards2")
	//card.print()

	card := newDeck()
	card.shuffle()
	card.print()

}

// func newCard() string {
//	return "Five Of Diamonds"
// }
