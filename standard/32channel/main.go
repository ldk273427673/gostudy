package main

import "fmt"

func main() {
	pipeline := make(chan int, 10)
	fmt.Printf("pipeline的容量是:%v\n", cap(pipeline))
	fmt.Println("开始往通道发送一个值")
	pipeline <- 1
	fmt.Printf("现在通道pipeline的容量是:%v\n", cap(pipeline))
	fmt.Printf("现在通道pipeline中有的数据量是:%v\n", len(pipeline))
}
