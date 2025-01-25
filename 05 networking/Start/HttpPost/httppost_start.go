package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postRequestTest() {
	const httpbin = "https://httpbin.org/post"

	// TODO: POST operation using Post
	reqBody := strings.NewReader(`
	{
		"Car Brand" : "Mercedes Benz",
		"Owner" : "Ahmed El marrouni"
	}
	`)

	res, err := http.Post(httpbin, "application/json", reqBody)

	if err != nil {
		fmt.Println(err)
		return
	}

	content, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Printf("%s\n", content)

	// TODO: POST operation using PostForm
	data := url.Values{}
	data.Add("firstname", "Ahmed")
	data.Add("lastname", "El marrouni")

	resp, err := http.PostForm(httpbin, data)
	content, _ = ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	fmt.Printf("%s\n", content)
}

func main() {
	// Execute a POST
	postRequestTest()
}
