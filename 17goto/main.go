package main

import "fmt"

func main() {
	goto test
	// fmt.Println("a")
test:
	fmt.Println("b")
}
