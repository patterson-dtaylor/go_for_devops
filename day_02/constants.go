package main

import "fmt"

const pi = 3.14

func main() {
	// pi = 400 -- this will error due to pi being a constant
	fmt.Println("PI = ", pi)
}
