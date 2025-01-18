package main

import "fmt"

func main() {
	// TODO: Print a simple string, no newline

	// TODO: Print with a newline

	// TODO: Print a string with values
	const answer = 42

	// TODO: Print a string with multiple interspersed values
	const a, b, c = 5, 5, 10

	// TODO: print a slice of data
	items := []int{10, 20, 40, 80}
	lenghth, err := fmt.Println(items)
	fmt.Println(lenghth, err)
}
