package main

import "fmt"

type Brand string

const
(
FACEBOOK Brand = "Facebook"
GOOGLE Brand = "Google"
)

func printBrand(brand Brand)  {
	fmt.Println("brand: ", brand);
}

func main()  {
	printBrand(FACEBOOK);
	printBrand(GOOGLE);
}