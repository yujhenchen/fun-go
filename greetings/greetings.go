package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// This function takes a name parameter whose type is string. The function also returns a string.
//
//	In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name.
func Hello(name string) (string, error) {
	// the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type)
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Note that randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).
func randomFormat() string {
	// declare a formats slice with three message formats. When declaring a slice, you omit its size in the brackets,
	//  like this: []string. This tells Go that the size of the array underlying the slice can be dynamically changed.
	formats := []string{
		"Hi, %v",
		"Good morning, %v",
		"Good night, %v",
	}
	return formats[rand.Intn(len(formats))]
}
