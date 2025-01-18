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

func decodeExample() {
	// TODO: declare some sample JSON data to decode
	// data := []byte(`
	// 	{
	// 		"fullname" : "John Q Public",
	// 		"addr" : "987 Main St",
	// 		"age": 45,
	// 		"favecolors" : ["Purple","White","Gold"]
	// 	}
	// `)

	// TODO: JSON will be decoded into a person struct

	// TODO: test to see if the JSON is valid, and then decode

	// TODO: data can also be decoded into a map structure

	// TODO: Unmarshal into a map

}

func main() {
	// Decode JSON into Go structs
	decodeExample()
}
