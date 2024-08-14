package main

import "fmt"

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards")

	hand, remainingCards := deal(cards, 5)

	// hand.print()
	// fmt.Println("Remaining:")
	remainingCards.print()

	fmt.Println(hand.toString())
}
