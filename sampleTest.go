package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	n := 15
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}

	}
	json.Valid()
}
