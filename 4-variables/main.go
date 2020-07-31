package main

import "fmt"


func main()  {

	// string variables
	
	//var message string = "Hello World"

	/*
	var message string 
	message = "Hello World"
	*/

	var message = "Hello World"

	fmt.Println(message)


	// int variables

	var a,b,c int

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var x,y,z int = 3, 8, 14

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)


	// define different variable types with single line
	var k,l,m = "k is a String variable", true, 5.7

	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)

	// defines variable shortly
	brand := "arjoy.net"
	fmt.Println(brand)

	v, n := true, 12.53
	fmt.Println(v, n)

	var myFloat32 float32 = 65.
	myComplex := complex(3, 5)

	println(myFloat32)
	println(myComplex)


	// defines constant variable
	const pi = 3.14
	println(pi)

	var (
		languageName = "GoLang"
		isGolangFast = true
		
	)

	
	println(languageName, isGolangFast)






}