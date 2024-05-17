package main

import "fmt"

var c int = 4

func main() {
	// || and && are logical operators the || is an or and the && is an and

	// You can use if else statement like normal

	if 7*2 != 14 {
		fmt.Println("This is not 14")
	} else {
		fmt.Println("This is 14")
	}

	a := 4

	if a*c == 16 {
		fmt.Println("This equal 16")
	} else if a*c == 14 {
		println("Somethign is test")
	} else {
		fmt.Println("None of these triggered")
	}

}
