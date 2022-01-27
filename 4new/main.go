package main

import "fmt"

func main() {
	ptr := new(int)
	fmt.Printf("ptr的值是%v\n", ptr)
	fmt.Println("new会创建一个指定类型的匿名变量，返回这个变量的内存地址，如果想要看这个值的话，需要使用*ptr")
	fmt.Println("ptr保存的值是：", *ptr)
}
