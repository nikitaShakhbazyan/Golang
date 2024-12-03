package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds",newCard()}
	cards = append(cards, "IDK") //adding to the end
	

	// first syntax

	// for i:=0; i<len(cards); i++ {
	// 	fmt.Println(i,cards[i])
	// }

	// second syntax

	for i,card:=range cards {
		fmt.Println(i,card)
	}
}

func newCard()string {
	return "Five of Diamonds"
}
