package main

import (
	"fmt"
	"strings"
)

// import (
// 	"fmt"
// 	"strings"
// )

func main() {
	s := "The quick brown fox jumps over the lazy dog"

	// Basic string operations

	// TODO: Length of string
	fmt.Println(len(s))

	// TODO: iterate over each character
	for _, ch := range s {
		fmt.Print(string(ch), ",")
	}
	fmt.Println()

	// TODO: Using operators < > == !=
	fmt.Println("dog" < "cat")
	fmt.Println("cat" > "bird")
	fmt.Println("dog" == "Dog")
	fmt.Println("dog" != "Dog")

	// TODO: Comparing two strings
	resultA := strings.Compare("dog", "cat")
	fmt.Println(resultA)

	resultB := strings.Compare("Horse", "Horse")
	if resultB == 1 {
		fmt.Println("True")
	} else {
		fmt.Println("false")
	}

	// TODO: EqualFold tests using Unicode case-folding
	fmt.Println(strings.EqualFold("Hello", "hello"))

	// TODO: ToUpper, ToLower, Title
	s1 := strings.ToUpper(s)
	s2 := strings.ToLower(s)
	s3 := strings.ToTitle(s)

	fmt.Println("Upper", s1)
	fmt.Println("Lower", s2)
	fmt.Println("Title", s3)
}
