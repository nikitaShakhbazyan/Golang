package main

import "fmt"

var node, golang, javs bool

var array1 []int

func main () {
	fmt.Println("hi")
	var x int
	fmt.Println(x,node)

	for i := 0 ; i<10; i++{
		array1 = append(array1, i)
	}
	array1= append([]int{15}, array1...)
	fmt.Println(array1)
}