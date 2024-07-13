package main

import "fmt"

func main () {
	// var x = 23
	// var y = 10
	// b:="nick"
	// var name string = "NIKITA"
	// var (
	// 	a string ="nikitos"
	// 	d int = 23
	// 	c bool = true
	// )
	// fmt.Println(x + y,b)
	// fmt.Println(name)
	// fmt.Println(a,d,c)
	temperature()
}

func temperature() {
	var fahrenheit float64
	fmt.Println("Введите температуру в градусах Фаренгейта:")
	fmt.Scan(&fahrenheit)

	celsius := fahrenheitToCelsius(fahrenheit)
	fmt.Printf("Температура в градусах Цельсия: %.2f\n", celsius)

	celsiusExample := fahrenheitToCelsius(20)
	fmt.Printf("Пример: 20 градусов Фаренгейта = %.2f градусов Цельсия\n", celsiusExample)
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}