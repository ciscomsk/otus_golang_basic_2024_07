package main

import "fmt"

var some string = "Some"
var xz1 string

//xz2 := "zzz" // syntax error: non-declaration statement outside function body

func main() {
	fmt.Println("Hello, world!")

	myName := "Mikhail"
	fmt.Printf("myName: %s", myName)

	var surname string
	surname = "Kuznetsov"
	fmt.Println(surname)

	fmt.Println(some)
	//some = 1 // cannot use 1 (untyped int constant) as string value in assignment
}
