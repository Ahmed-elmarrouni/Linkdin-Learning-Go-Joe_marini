package main

import (
	"fmt"
	"math"
)

func main() {
	// TODO: print Pi
	fmt.Println(math.Pi)

	// TODO: Ceiling and Floor functions
	fmt.Println("Ceiling :", math.Ceil(math.Pi))
	fmt.Println("Floor :", math.Floor(math.Pi))

	// TODO: Truncate will return int value of X
	fmt.Printf("%.2f\n", math.Trunc(math.Pi))

	// TODO: Max and Min values
	fmt.Println("Max value", math.Max(math.Pi, math.Ln2))
	fmt.Println("Min value", math.Min(math.Pi, math.Ln2))

	// TODO: Mod operator is like % but for floats
	fmt.Println(17 % 5)
	fmt.Println(math.Mod(17.0, 5.0))

	// TODO: Round and RoundToEven
	fmt.Printf("%.1f\n", math.Round(10.5))
	fmt.Printf("%.1f\n", math.Round(-10.5))

	fmt.Printf("%.1f\n", math.RoundToEven(10.5))
	fmt.Printf("%.1f\n", math.RoundToEven(11.5))
}
