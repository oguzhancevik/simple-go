package main

import "fmt"

func main() {

	var myArr = []int{1, 2, 4, 8}

	for i, v := range myArr {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// second usage
	for i := range myArr {
		fmt.Println("Array item", i, "is", myArr[i])
	}

	// usage with map
	cities := map[string]string{"Turkey": "Istanbul", "France": "Paris", "Italy": "Roma"}
	for key := range cities {
		fmt.Println("City of", key, "is", cities[key])
	}

}
