package main

import "fmt"

// Создаем новый тип 'deck' — это слайс строк
type deck []string

// Создаем новый тип 'carding' — это строка
type carding string

// Метод для печати карт
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
