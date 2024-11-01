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

func Hellos(names []string) (map[string]string, error) {
	// initialize a map with the following syntax: make(map[key-type]value-type)
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	// = vs :=

	// = (Assignment Operator):
	// Used to assign a new value to an already declared variable.
	// Works with variables that have been previously defined with var or :=

	// := (Short Variable Declaration):
	// Used to declare and initialize a new variable in a concise way.
	// Can only be used within functions (not at the package level).
	// It both declares and assigns the variable in one step.

	return messages, nil
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
