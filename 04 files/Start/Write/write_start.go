package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

// func handleErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func checkFileExists(filePath string) bool {
// 	if _, err := os.Stat(filePath); err != nil {
// 		if os.IsNotExist(err) {
// 			return false
// 		}
// 	}
// 	return true
// }

func writeFileData() {
	// TODO: create a new file

	// TODO: it is idiomatic in Go to defer the close after you open the file

	// TODO: get the Name of the file

	// TODO: write some string data to the file

	// TODO: write some individual bytes to the file

	// TODO: the Sync function forces the data to be written out
}

func appendFileData(fname string, data string) {
	// TODO: Use the lower-level OpenFile function to open the file in append mode
}

func main() {
	// TODO: Simple example: dump some data to a file

	// TODO: More complex example: Granular writing of data

	// TODO: append data to a file
}
