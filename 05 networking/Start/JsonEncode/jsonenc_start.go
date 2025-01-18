package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

type person struct {
	Name       string
	Address    string
	Age        int
	FaveColors []string
}

func encodeExample() {
	// TODO: create some people data
	// people := []person{
	// 	{"Jane Doe", "123 Anywhere Street", 35, nil},
	// 	{"John Public", "456 Everywhere Blvd", 31, []string{"Purple", "Yellow", "Green"}},
	// }

	// TODO: Marshal is used to convert a data structure to JSON format
	// MarshalIndent is used to format the JSON string with indentation

}

func main() {
	// Encode Go data as JSON
	encodeExample()
}
