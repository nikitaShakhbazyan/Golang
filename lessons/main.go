package main

// import "fmt"

func main() {
	cards := deck{"Ace of Diamonds",newCard()}
	cards = append(cards, "IDK") //adding to the end
	
	cards.print()
}

func newCard()string {
	return "Five of Diamonds"
}
