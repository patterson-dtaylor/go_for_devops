package main

import "fmt"

var (
	foo  int    // 0
	bar  string //""
	fizz bool   // false
	buzz *int   // nil
)

func main() {
	fmt.Println("foo=", foo)
	fmt.Println("bar=", bar)
	fmt.Println("fizz=", fizz)
	fmt.Println("buzz=", buzz)
}
