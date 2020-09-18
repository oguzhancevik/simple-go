package main

import (
	"fmt"
)

func main() {

	// int array
	myIntArray := [3]int{}
	myIntArray[0] = 65
	myIntArray[1] = 34
	myIntArray[2] = 81
	fmt.Println(myIntArray)

	// string array
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[1])
	colors[1] = "White"
	fmt.Println(colors)
	fmt.Println(colors[1])

	// other usage
	var numbers = [5]int{1, 3, 5, 7, 9}
	fmt.Println("length of array", len(numbers))

	// other usage
	var mySpecialArray = [...]int{0, 2, 4, 6, 8}
	fmt.Println(mySpecialArray)

	//other usage
	var cars [3]string
	cars[0] = "Tesla"
	cars[1] = "Mercedes"
	cars[2] = "BMW"
	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}
	for index, value := range cars {
		fmt.Println("index:", index, " value:", value)
	}

}
