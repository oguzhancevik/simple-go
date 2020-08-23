package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	input1, _ := reader.ReadString('\n')
	fmt.Println("str: " + input1)

	fmt.Println("Enter a number: ")
	input2, _ := reader.ReadString('\n')
	number, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Number: ", number)
	}
}
