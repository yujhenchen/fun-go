package greetings

import (
	"fmt"
)

// This function takes a name parameter whose type is string. The function also returns a string.
//
//	In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name.
func Hello(name string) string {
	// the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type)
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}
