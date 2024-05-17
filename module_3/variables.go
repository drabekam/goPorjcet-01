package main

import "fmt"

func main() {

	// Go does not allow you to create an empty variable it will throw an error if you do this

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // var can be used to make multiple functions that are needed this can be inside and outside a function
	fmt.Println(e)

	f := "apple" // This is the short hand for variables and it only able to created like this in functions
	fmt.Println(f)
}
