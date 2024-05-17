package main

import (
	"fmt"
	"math"
)

const s string = "constant_test"

func main() {
	// const are things that can not change. They are the same throughout the lifecycle of the application

	// you declare them in the same manner as variable but with const you can not use short hand for this

	const n = 5000000

	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
