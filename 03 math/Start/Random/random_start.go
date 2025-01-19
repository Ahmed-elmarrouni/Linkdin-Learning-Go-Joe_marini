package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// TODO: Initialize the random seed to an unknown value
	rand.Seed(time.Now().UnixNano())

	// TODO: generate random Integer numbers
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(10))

	// TODO: generate random Floating Point numbers
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())

	// TODO: use the Perm function to create permutations
	s := []string{"apple", "pear", "grape", "orange", "kiwi", "melon"}
	indexes := rand.Perm(len(s))
	fmt.Println("Index value :", indexes)

	for i := 0; i < len(indexes); i++ {
		fmt.Println(s[indexes[i]])
	}

	// TODO: Generate a random integer between a and b:
	const a, b = 10, 50

	for i := 0; i < 5; i++ {
		num := a + rand.Intn(b-a+1)
		fmt.Println(num, " ")
	}
	fmt.Println()

	// TODO: Generate a random uppercase character:
	for i := 0; i < 10; i++ {
		char := string('A' + rune(rand.Intn('Z'-'A'+1)))
		fmt.Println(char, " ")
	}
}
