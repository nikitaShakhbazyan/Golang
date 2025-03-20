package main

import (
	"fmt"
	"math"
)


func main(){
	 userHeight := 1.8
	var userKG float64 = 100 // can't use userKG float64 := 100 - point a type needs to be declared as VAR or
	// var userKG float64
	// userKG = 100  OR
	// var userHeight,userKG float64 = 1.8,100
	var IMT = userKG / math.Pow(userHeight,2)

	fmt.Print(IMT)
}