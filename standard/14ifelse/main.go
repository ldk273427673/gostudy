package main

import "fmt"

func main() {
	age := 20
	if age > 20 {
		fmt.Println("你成年了")
	} else if age == 20 {
		fmt.Println("你正好20岁")
	} else {
		fmt.Println("你还小")
	}
}
