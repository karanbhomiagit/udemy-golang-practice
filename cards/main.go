package main

func main() {
	//var card string = "Ace of spades"
	// var card = "Ace of spades"

	// card := "Ace of spades"
	// fmt.Println(card)
	// card = newCard()

	// cards := []string{"Ace of Spades", newCard()}
	// cards = append(cards, "Six of Spades")
	// fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// 	fmt.Println(card)
	// }

	// cards := newDeck()
	// hand, cards := deal(cards, 4)
	// hand.print()
	// cards.print()

	cards1 := newDeckFromFile("abc.txt")
	cards1.print()

	cards1.shuffle()
	cards1.print()

	cards1.saveToFile("abc.txt")
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
