package main

import (
	"fmt"
	"unsafe"
)

type nullStruct struct{}

func main() {
	l := nullStruct{}
	fmt.Println("空结构体可以不占空间，所以输出的结果会是0")
	fmt.Println(unsafe.Sizeof(l))
}
