package main

import "fmt"

func main() {
	str1 := "programming language"
	str2 := "GO"
	str3 := "Created by Google"
	number := 1996
	isTrue := true

	stringLength, _ := fmt.Println(str1, str2, str3)

	fmt.Println("String length: ", stringLength)

	fmt.Printf("Value of number: %v\n", number)

	fmt.Printf("Value of isTrue: %v\n", isTrue)

	fmt.Printf("Value of number as float: %.2f\n", float64(number))

	fmt.Printf("Data Types: %T, %T, %T, %T, and %T\n", str1, str2, str3, number, isTrue)

	myStr := fmt.Sprintf("Data Types: %T, %T, %T, %T, and %T\n", str1, str2, str3, number, isTrue)
	fmt.Println("mystr: " + myStr)

}
