package main

import "fmt"

func main() {

	myArray := [...]int{10, 20, 30}
	mySlice := myArray[:]
	fmt.Println(mySlice)
	mySlice[0] = 90
	fmt.Println(mySlice)
	fmt.Println(myArray)

	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Black", "White")
	fmt.Println(colors)
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 48
	numbers[1] = 274
	numbers[2] = 368
	numbers[3] = 18
	numbers[4] = 73
	fmt.Println(numbers)
	numbers = append(numbers, 346581)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))
}
