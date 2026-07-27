package core

import "fmt"

// IfThenElse returns the valueIfTrue if the condition is true, otherwise it returns the valueIfFalse
func IfThenElse[T any](condition bool, valueIfTrue, valueIfFalse T) T {
	if condition {
		return valueIfTrue
	}
	return valueIfFalse
}

// Println prints the given message to the console, formatting it with the given arguments if any
func Println(msg string, args ...any) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	fmt.Println(msg)
}
