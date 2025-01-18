package main

import (
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// TODO: Shuffle a sequence of values
	const numstring = "one two three four five six"

	// TODO: Generate a random string
	const upletters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const lowletters = "abcdefghijklmnopqrstuvwxyz"
	const digits = "0123456789"
	const specialchars = "~=+%^*()[]{}!@#$?|"
	const allchars = upletters + lowletters + digits + specialchars

	// TODO: Generate random string with rules
}
