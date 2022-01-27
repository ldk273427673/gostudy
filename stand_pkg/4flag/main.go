package main

import (
	"flag"
	"fmt"
)

var canshu string
var debug bool

func init() {
	flag.StringVar(&canshu, "n", "ls", "input the canshu")
	flag.BoolVar(&debug, "debug", false, "是否开启 DEBUG 模式")
}

func main() {
	// var name string
	// flag.StringVar(&name, "n", "dengke", "input you name")
	// flag.Parse()
	// fmt.Printf("你输入的name参数是%v\n", name)
	flag.Parse()
	fmt.Println(canshu)
	fmt.Println(debug)
}
