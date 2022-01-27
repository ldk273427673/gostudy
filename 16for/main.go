package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
	var count int = 1
	for {
		if count > 10 {
			break
		}
		fmt.Println(count)
		count++
	}

}
