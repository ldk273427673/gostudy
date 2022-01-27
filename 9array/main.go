package main

import "fmt"

func main() {
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Printf("arr1: %v\n", arr1)

	arr2 := [3]int{1, 2, 3}
	fmt.Printf("arr2: %v\n", arr2)

	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arr3: %v\n", arr3)

	type arrtype [7]int
	myarry := arrtype{1, 2, 3, 4, 5, 5, 6}
	fmt.Printf("myarry: %v\n", myarry)
}
