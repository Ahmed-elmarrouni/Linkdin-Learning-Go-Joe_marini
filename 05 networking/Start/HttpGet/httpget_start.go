package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getRequestTest() {
	const httpbin = "https://httpbin.org/get"

	// TODO: Perform a GET operation
	resp, err := http.Get(httpbin)
	if err != nil {
		return
	}

	// TODO: The caller is responsible for closing the response
	defer resp.Body.Close()

	// TODO: We can access parts of the response to get information:
	fmt.Println("Status : ", resp.Status)
	fmt.Println("Status code: ", resp.StatusCode)
	fmt.Println("Protocol : ", resp.Proto)
	fmt.Println("Content lenght : ", resp.ContentLength)

	// TODO: Use a String Builder to build the content from bytes
	var stringBuilder strings.Builder

	// TODO: Read the content and write it to the builder
	content, _ := ioutil.ReadAll(resp.Body)
	byteCount, _ := stringBuilder.Write(content)

	// TODO: Format the output
	fmt.Println(byteCount, stringBuilder.String())

}

func main() {
	// Execute a GET request
	getRequestTest()
}
