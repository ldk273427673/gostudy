package main

import "fmt"

func main() {
	var age int = 20
	var ptr = &age
	fmt.Printf("age的值是%d\n", age)
	fmt.Printf("ptr的值是%v\n", ptr)
}
