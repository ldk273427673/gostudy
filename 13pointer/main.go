package main

import "fmt"

func main() {
	var name string = "张飞"
	ptrname := &name
	fmt.Printf("name: %v\n", name)
	fmt.Printf("ptrname: %v\n", ptrname)
	fmt.Println("指针ptrname中存储的name值是:", *ptrname)

	fmt.Println("第一种创建指针的方法是有个变量，然后直接使用变量的内存地址来赋予这个指针变量")
	fmt.Println("第二种创建指针的方法是使用new")
	ptr1 := new(string)
	fmt.Printf("ptr1的类型是%T", ptr1)

	a := 25
	var b *int // 声明一个指针

	if b == nil {
		fmt.Println(b)
		b = &a // 初始化：将a的内存地址给b
		fmt.Println(b)
	}
}
