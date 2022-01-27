package main

import "fmt"

func main() {
	fmt.Println(`
	int 有符号，根据系统是32位还是64位，宽度是32位bit或者64位bit，占用4/8个字节
	int8 有符号,宽度8bit，占用1个字节
	int16 有符号，宽度16bit，占用2个字节
	int32 有符号，宽度32bit，占用4个字节
	int64 有符号，宽度16bit，占用8个字节
	uint 无符号，根据系统是32位还是64位，宽度是32位bit或者64位bit，占用4/8个字节
	uint8 无符号,宽度8bit，占用1个字节
	uint16 无符号，宽度16bit，占用2个字节
	uint32 无符号，宽度32bit，占用4个字节
	uint64 无符号，宽度16bit，占用8个字节
	`)
	var num1 int = 0b11111
	var num2 int = 0o121212
	var num3 int = 0xcc
	fmt.Println(num1, num2, num3)
}
