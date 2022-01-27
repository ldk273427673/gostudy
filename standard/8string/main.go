package main

import "fmt"

func main() {
	str1 := "hello"
	fmt.Printf("str1: %v\n", str1)
	str2 := "nihao,邓肯"
	fmt.Println(len(str2))
	fmt.Println("字符串中，使用utf8编码，英文占用1个字节也就是8位，中文占用3个字节，所以这个nihao，邓肯，使用的字节数是5+1+3*2=12")
}
