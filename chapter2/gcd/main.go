package main

import (
	"flag"
	"fmt"
)

func main() {
	var first int
	var second int

	flag.IntVar(&first, "first", 0, "First number")
	flag.IntVar(&second, "second", 0, "Second number")

	flag.Parse()

	fmt.Println("First number: ", first)
	fmt.Println("Second number: ", second)

	for second != 0 {
		rem := first % second

		first = second
		second = rem
	}

	fmt.Println("GCD is", first)
}
