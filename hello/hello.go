package main

import (
	"fmt"
	"log"

	"greetings"
)

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Println(quote.Go())

	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)
	// message, err := greetings.Hello("test")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(message)
	names := []string{"A", "B", "C"}
	msgs, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msgs)
}
