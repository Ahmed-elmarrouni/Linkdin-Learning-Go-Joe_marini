package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	strVar, _ := reader.ReadString('\n')
	fmt.Println("Your name is:", strVar)
}
