package main

import (
	"fmt"
	"strings"
)

func main() {
	// TODO: Declare an empty strings.Builder
	var sb strings.Builder

	// TODO: Write some content
	sb.WriteString("This is a string 1 \n")
	sb.WriteString("This is a string 2 \n")
	sb.WriteString("This is a string 3 \n")

	// TODO: Output the concatenated string
	fmt.Println(sb.String())

	// TODO: Examine the builder's capacity
	fmt.Println("Capacity:", sb.Cap())

	// TODO:
	// Grow the capacity - use this if you know in advance how much data
	// you're going to be writing into the buffer to minimize copies
	sb.Grow(1024)
	fmt.Println("Capacity:", sb.Cap())

	for i := 0; i <= 10; i++ {
		fmt.Println(&sb, "string %d --", i)
	}

	fmt.Println(sb.String())

	// TODO: we can get the length of what the final string will be
	fmt.Println("Builder size is", sb.Len())

	// TODO: The Reset function will reset the builder to original state
	sb.Reset()
	fmt.Println("After reset :")
	fmt.Println("Capacity:", sb.Cap())
	fmt.Println("Builder size is", sb.Len())

}
