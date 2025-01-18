package main

import "fmt"

type circle struct {
	radius int
	border int
}

func main() {
	x := 20
	f := 123.45

	// Basic formatting
	fmt.Printf("%d\n", x)
	fmt.Printf("%x\n", x)

	// Booleans can be printed as "true" or "false"
	fmt.Printf("%t\n", x > 10)

	// floating point numbers
	fmt.Printf("%f\n", f)
	fmt.Printf("%e\n", f)

	// Using explicit argument indexes
	fmt.Printf("%[2]d %[1]d\n", 52, 40)
	// Argument indexes can be used to print values repeatedly
	fmt.Printf("%d %#[1]o %#[1]x\n", 52)

	// Print a value in default format
	c := circle{
		radius: 20,
		border: 5,
	}
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%T\n", c)

	// Sprintf is the same as Printf, but returns a string
	s := fmt.Sprintf("%[2]d %[1]d\n", 52, 40)
	fmt.Println(s)
}
