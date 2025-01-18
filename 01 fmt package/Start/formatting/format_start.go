package main

import "fmt"

//import "fmt"

func main() {
	f := 123.4567

	// TODO: control the decimal precision
	fmt.Printf("%.2f\n", f)

	// TODO: print with width 10 and default precision
	fmt.Printf("%10f\n", f)

	// TODO: print with padding and precision
	fmt.Printf("%10.2f\n", f)

	// TODO: always use a + sign
	fmt.Printf("%+10.2f\n", f)

	// TODO: pad with 0s instead of spaces
	fmt.Printf("%010.2f\n", f)

}
