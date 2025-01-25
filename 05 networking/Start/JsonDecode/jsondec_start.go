package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string // `json:"fullname"`
	Address    string // `json:"addr"`
	Age        int
	FaveColors []string
}

func decodeExample() {
	// TODO: declare some sample JSON data to decode
	data := []byte(`
		{
			"fullname" : "John Q Public",
			"addr" : "987 Main St",
			"age": 45,
			"favecolors" : ["Purple","White","Gold"]
		}
	`)

	// TODO: JSON will be decoded into a person struct
	var someone person

	// TODO: test to see if the JSON is valid, and then decode
	valid := json.Valid(data)
	if valid {
		json.Unmarshal(data, &someone)
		fmt.Printf("%#v\n", someone)
	}

	// TODO: data can also be decoded into a map structure
	var m map[string]interface{}

	// TODO: Unmarshal into a map
	json.Unmarshal(data, &m)
	fmt.Printf("%#v\n", m)

	for k, v := range m {
		fmt.Printf("Key: %v, Value: %v (%T)\n", k, v, v)
	}
}

func main() {
	// Decode JSON into Go structs
	decodeExample()
}
