package main

import "fmt"

func main () {
	// fmt.Println(1)
	// fmt.Println(2)
	// fmt.Println(3)
	// fmt.Println(4)
	// fmt.Println(5)
	// fmt.Println(5)
	// fmt.Println(6)
	// fmt.Println(7)
	// fmt.Println(8)
	// fmt.Println(9)
	// fmt.Println(10)

	// i := 1

	// for i<=10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }
	// loop


    array := [6]int{1, 2, 3, 4, 5, 6,}
	
	// for i:=1 ; i<= len(array); i++{
	// 	if i % 2 == 0{
	// 		fmt.Println(i,"EVEN")
	// 	} else {
	// 		fmt.Println(i,"odd")
	// 	}
	// }

	for i := 0; i < len(array); i++ {
        switch array[i] {
        case 0:
            fmt.Println("Zero")
        case 1:
            fmt.Println("One")
        case 2:
            fmt.Println("Two")
        case 3:
            fmt.Println("Three")
		case 4: 
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		case 6:
			fmt.Println("Six")
        default:
            fmt.Println("Other number")
        }
    }
}