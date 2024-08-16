package main

func main() {
	// cards := newDeck()
	cards := newDeckFromFile("my_cards")

	// cards.saveToFile("my_cards")

	// hand, remainingCards := deal(cards, 5)

	cards.shuffle()

	cards.print()
	// // fmt.Println("Remaining:")
	// remainingCards.print()

	// fmt.Println(hand.toString())
}
