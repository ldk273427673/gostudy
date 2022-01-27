package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		for i, j := range os.Args {
			fmt.Printf("参数%v的值是%v\n", i, j)
		}
	}
}
