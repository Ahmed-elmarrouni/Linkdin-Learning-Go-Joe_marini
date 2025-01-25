package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Define a URL
	s := "https://www.example.com:8000/user?username=joemarini"

	// TODO: Use the Parse function to parse the URL content
	result, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Past: ", result.Path)
	fmt.Println("Port: ", result.Port())
	fmt.Println("RawQuery: ", result.RawQuery)

	// TODO: Extract the query components into a Values struct
	vals := result.Query()
	fmt.Println(vals["username"])

	// TODO: create a URL from components
	newurl := &url.URL{
		Scheme:   "https",
		Host:     "www.ahmed.com",
		Path:     "/info",
		RawQuery: "id=123",
	}
	s = newurl.String()
	fmt.Println(s)

	newurl.Host = "www.elmarrouni.com"
	s = newurl.String()
	fmt.Println(s)

	// TODO: create a new Values struct and encode it as a query string
	newVal := &url.Values{}
	newVal.Add("x", "123")
	newVal.Add("y", "Mercedess")
	newurl.RawQuery = newVal.Encode()
	s = newurl.String()
	fmt.Println(s)

}
