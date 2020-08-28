package main

import "fmt"

func main() {

	// break
	i := 0
	for {
		if i > 3 {
			break
		}
		fmt.Println("value of i: ", i)
		i++
	}

	// continue
	fmt.Println("odd numbers between 1 and 10")
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}

}
