package main

import (
	f "fmt"
	_ "image/png"
)

func main() {
	f.Println("重命名包名字，比如用f代替fmt，可以直接用f来调用这个包")
	f.Println("hello")
	f.Println("使用_符号，可以导入一个包不用，不会影响编译报错")
}
