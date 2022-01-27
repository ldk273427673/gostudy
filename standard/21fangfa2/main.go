package main

import "fmt"

type Profile struct {
	name           string
	age            int
	gender         string
	mother, father *string
}

func (p Profile) p_fangfa() {
	fmt.Printf("名字是: %v\n", p.name)
	fmt.Printf("年龄是: %v\n", p.age)
	fmt.Printf("性别是: %v\n", p.gender)

}

func (p *Profile) increace_age() {
	p.age += 2
}
func main() {
	p1 := Profile{
		name:   "张飞",
		age:    20,
		gender: "男",
	}
	fmt.Printf("p1: %v\n", p1)

	fmt.Println("调用方法")
	p1.p_fangfa()

	fmt.Println("调用增加年龄的方法")
	p1.increace_age()
	fmt.Printf("p1.age: %v\n", p1.age)

}
