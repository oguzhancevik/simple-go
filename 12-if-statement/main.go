package main

import "fmt"

func main() {
	a := 10
	b := 10

	if a > b {
		fmt.Println("a bigger than b")
	} else if a == b {
		fmt.Println("a equals b")
	} else {
		fmt.Println("a smaller than b")
	}

	if foo := 5; foo > 3 {
		fmt.Println("foo is big than 3")
	}

	// we can not use foo. because foo is only use if scope
	// fmt.Println(foo)

}
