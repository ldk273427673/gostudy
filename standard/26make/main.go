package main

import "fmt"

func main() {
	make1 := make([]int, 2, 10)
	fmt.Printf("make1: %v\n", make1)
	fmt.Printf("make1的类型是%T", make1)
	fmt.Println("make只适用于切片，map和管道，用来给这三种类型的变量分配内存地址，并初始化，返回的是这个类型的本身，比如切片返回的就是切片类型")
}
