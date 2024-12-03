package main

import "fmt"

func main() {
	var card string = "Ace of Spades"
	// type can be declared in this case like this
	card = "DIDID"
	var example float64 = 3132.22
	var integer int = 2003
	var logic bool = false
	fmt.Println(card)
	fmt.Println(example)
	fmt.Println(integer)
	fmt.Println(logic)

	card2 := "IDK"
	// and here we can write value without declaring type but if we write another type to the variable there will be an error
	// card2 = false
	fmt.Println(card2)

}
