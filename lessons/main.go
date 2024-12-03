package main

// import "fmt"

func main() {
	// Приводим результат newCard() к типу string
	cards := deck{"Ace of Diamonds", string(newCard())}
	cards = append(cards, "IDK") 
	cards.print()
}

// Функция для создания новой карты, возвращающая тип carding
func newCard() carding {
	return carding("Five of Diamonds")
}
