package main

import (
	"fmt"
	"os"
)

func main() {
	uName := os.Getenv("USERNAME")
	uDomain := os.Getenv("USERDOMAIN")
	processorArchitecture := os.Getenv("PROCESSOR_ARCHITECTURE")
	processorIdentifier := os.Getenv("PROCESSOR_IDENTIFIER")
	processorLevel := os.Getenv("PROCESSOR_LEVEL")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")

	fmt.Println("uName: " + uName)
	fmt.Println("uDomain: " + uDomain)
	fmt.Println("processorArchitecture: " + processorArchitecture)
	fmt.Println("processorIdentifier: " + processorIdentifier)
	fmt.Println("processorLevel: " + processorLevel)
	fmt.Println("goRoot: " + goRoot)
	fmt.Println("goPath: " + goPath)
	fmt.Println("homePath: " + homePath)
}
