package main

import (
	"fmt"
	"math"
)


func main(){
	var userHeight = 1.8
	var userKG float64 = 100
	var IMT = userKG / math.Pow(userHeight,2)

	fmt.Print(IMT)
}