package main

import "fmt"

func main() {
	fmt.Println("hello")
	n := 104
	fmt.Printf("n的二进制表示是%b\n", n)
	fmt.Printf("n的八进制表示是%o\n", n)
	fmt.Printf("n的十六进制表示是%d\n", n)

	fmt.Printf("n的默认显示值是%v\n", n)
	fmt.Printf("n的类型是%T\n", n)

	f := 1.024
	fmt.Printf("%9.2f\n", f)

}
