package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// create a directory
	os.Mkdir("newdir", os.ModePerm)
	// Remove will remove an item
	defer os.Remove("newdir")

	// create a deep directory
	os.MkdirAll("path/to/new/dir", os.ModePerm)
	// RemoveAll will remove an item and all children
	defer os.RemoveAll("path")

	// get the current working directory
	dir, _ := os.Getwd()
	fmt.Println("Current dir is:", dir)

	// get the directory of the current process
	exedir, _ := os.Executable()
	fmt.Println("Current dir is:", exedir)

	// read the contents of a directory
	contents, _ := ioutil.ReadDir(".")
	for _, fi := range contents {
		fmt.Println(fi.Name(), fi.IsDir())
	}
}
