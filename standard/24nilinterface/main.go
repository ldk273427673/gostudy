package main

import (
	"fmt"
)

func main() {
	var i interface{}
	fmt.Printf("type:%T,value:%v\n", i, i)

	i = 1
	fmt.Printf("i: %v\n", i)

	i = "dengken"
	fmt.Printf("i: %v\n", i)

	i = true
	fmt.Printf("i: %v\n", i)

	fmt.Println("实例：函数接受一个任意值的参数")
	a := 1
	anyVar(a)
	anyVar("sdd")

	fmt.Println("实例：函数接受任意个，任意类型参数")
	anyVar2(2, 2, 3, "dfdfd", '2', true)
}

func anyVar(v interface{}) {
	fmt.Printf("v: %v\n", v)
}
func anyVar2(v ...interface{}) {
	for i, j := range v {
		fmt.Printf("第%v个值是%v\n", i, j)
	}
}
