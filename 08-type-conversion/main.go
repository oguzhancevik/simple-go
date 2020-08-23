package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*
		CONVERSION
	*/

	// conversion string to int
	var numberString = "1453"
	number, _ := strconv.Atoi(numberString)
	result := number + 8
	fmt.Println(result)

	// conversion int to string
	var numeric = 5
	numericStr := strconv.Itoa(numeric)
	fmt.Println("numericStr: " + numericStr)

	/*
		CASTING
	*/

	var i int = 1071
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, f, u)

}
