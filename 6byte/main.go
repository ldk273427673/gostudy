package main

import "fmt"

func main() {
	fmt.Println("byte等同于无符号的8位数，等价于uint8")
	fmt.Println("其实对应的是ascii码的码值")
	var a byte = 'A'
	var b uint8 = 'B'
	// var c byte = '中'
	fmt.Printf("b: %v\n", b)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("c的值如果是中文，就打印不出来，因为byte表示的范围有限，有这个2的8次方，中文是unicode编码，表示不出来，需要用rune")

}
