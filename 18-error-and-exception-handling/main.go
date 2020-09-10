package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	// sample 1
	file, err := os.Open("invalid-file-path")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
	}

	// sample 2
	mySpecialError := errors.New("this is my special error")
	fmt.Println(mySpecialError)

	// sample 3
	grade := -5
	if grade < 0 {
		fmt.Println(fmt.Errorf("this grade is too bad"))
	}
}
