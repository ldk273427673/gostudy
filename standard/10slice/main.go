package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)
	fmt.Printf("slice1的类型是%T\n", slice1)

	fmt.Println("使用make函数来构造切片")
	makeslice := make([]int, 2, 5)
	fmt.Printf("makeslice: %v\n", makeslice)
	fmt.Println(len(makeslice))
	fmt.Println(cap(makeslice))

	var myarr []int
	fmt.Println(myarr == nil)
}
