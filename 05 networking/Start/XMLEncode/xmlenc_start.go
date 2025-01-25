package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	XMLName    xml.Name `xml:"personData"`
	FirstName  string   `xml:"first_name"`
	LastName   string   `xml:"last_name"`
	Address    string   `xml:"address"`
	Age        int      `xml:"age"`
	FaveColors []string `xml:"fave_colors>color"`
}

func main() {
	// TODO: Declare some sample data
	people := []person{
		{FirstName: "Jane", LastName: "Doe", Address: "123 Anywhere Street", Age: 35},
		{FirstName: "John", LastName: "Public", Address: "456 Everywhere Blvd", Age: 29,
			FaveColors: []string{"Purple", "Yellow", "Green"}},
	}

	// TODO: The MarhsalIndent function indents the XML text
	// result, err := xml.Marshal(people)
	result, err := xml.MarshalIndent(people, "", "\t")
	if err != nil {
		return
	}
	fmt.Printf("%s\n", result)
}
