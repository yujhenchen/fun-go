package main

import (
	"fmt"

	"greetings"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	message := greetings.Hello("Dog")
	fmt.Println(message)
}
