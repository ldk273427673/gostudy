package main

import "fmt"

func sum(args ...int) int {
	var sum int
	for _, v := range args {
		sum += v
	}
	return sum
}

func twoResult(a, b int) (c, d int) {
	c = a + b
	d = a - b
	return c, d
}
func main() {
	fmt.Println(sum(1, 2, 3, 4))
	i, j := twoResult(1, 3)
	fmt.Println(i, j)
}
