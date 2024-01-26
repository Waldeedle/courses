package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	// fmt.Println(cards)
}

// func main() {
// 	// var card string = "Ace of Spades"
// 	// card := "Ace of Spades"
// 	// card = "Five of Diamonds"
// 	card := newCard()

// 	fmt.Println(card)
// }

// func newCard() int {
// 	return 12
// }

func newCard() string {
	return "Five of Diamonds"
}
