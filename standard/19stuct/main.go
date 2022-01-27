package main

import "fmt"

type Profile struct {
	name           string
	age            int
	gender         string
	mother, father *string
}

func main() {
	p1 := Profile{
		name:   "张飞",
		age:    20,
		gender: "男",
	}
	fmt.Printf("p1: %v\n", p1)

}
