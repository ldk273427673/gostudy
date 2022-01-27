package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	fmt.Println("new是内置函数，需要一个类型参数，可以是系统自带，也可以是自定义的类型，会分配一个内存地址，然后初始化一个空值，并返回一个指针")
	new1 := new(int)
	fmt.Printf("new1的内存地址是: %v\n", new1)  //0xc000014098
	fmt.Printf("new1的初始化值是: %v\n", *new1) //0

	fmt.Println("当用一个自定义类型new的话")
	new2 := new(Student)
	fmt.Printf("new2: %v\n", new2)
	fmt.Printf("new2: %v\n", *new2)
}
