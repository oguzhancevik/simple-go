package main

import "fmt"

func main()  {
	
	a :=10;
	b := 10;

	total := a+b;
	fmt.Println("total: ", total);

	total = total -5;
	fmt.Println("total: ", total);

	total *= 2;
	fmt.Println("total: ", total);

	total = total / 5;
	fmt.Println("total: ", total);

	total = total % 5;
	fmt.Println("total: ", total);

	total++;
	fmt.Println("total: ", total);
}