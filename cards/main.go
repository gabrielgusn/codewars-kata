package main

import "fmt"

func main() {

	cards := deck{}
	cards = cards.newDeck()

	hand, _ := deal(cards, 5)

	hand.printDeck()
	// table.printDeck()
	hand.shuffle()
	fmt.Println()
	hand.printDeck()

	// err := cards.saveDeckToFile("my_cards")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(newDeckFromFile("teste.txt"))

	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// newDeck := newDeckFromFile("teste.txt")

	// fmt.Println(newDeck.toString())
}
