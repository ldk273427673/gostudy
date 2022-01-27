package main

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

func main() {
	fmt.Println("Hello,", Magenta("Aurora"))
	fmt.Println(Bold(Cyan("Cya!")))
	fmt.Println(Sprintf(Magenta("Got it %d times"), Green(1240)))
	r := Red("red")
	var i int
	fmt.Printf("%T %p\n", r, Green(&i))
}
