package main

import "fmt"

func main() {
	s1 := []int{1}
	fmt.Printf("s1: %v\n", s1)

	fmt.Println("追加一个元素")
	s1 = append(s1, 2)
	fmt.Printf("s1: %v\n", s1)

	fmt.Println("追加多个元素")
	s1 = append(s1, 3, 4)
	fmt.Printf("s1: %v\n", s1)

	fmt.Println("追加一个切片")
	s1 = append(s1, []int{6, 7, 8, 9}...)
	fmt.Printf("s1: %v\n", s1)
}
