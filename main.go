package main

func main() {
	//var card string = "Ace of Spade"
	//card := newCard()
	//cards := deck{"Ace of Spades", newCard()}
	//cards = append(cards, "Two of Spades") //returns a new slice
	//cards := newDeck()
	//cards.print()
	//hand, remaingDeck := deal(cards, 4)
	//hand.print()
	//remaingDeck.print()
	//fmt.Println(hand.toString())
	//cards.savetoFile("AllCards")
	//fileDeck := newDeckFromFile("AllCards")
	//fileDeck.print()
	cards := newDeck()
	cards.shuffelDeck()
	cards.print()

}
