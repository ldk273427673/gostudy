package main

import "fmt"

func main() {
	map1 := map[string]int{"zhangfei": 18, "zhaoyun": 20}
	fmt.Printf("map1: %v\n", map1)

	map2 := make(map[string]int)
	map2["刘备"] = 30
	fmt.Printf("map2: %v\n", map2)

	delete(map1, "zhaoyun")
	fmt.Printf("map1: %v\n", map1)

	name, ok := map1["zhaoyun"]
	if ok {
		fmt.Println("赵云的年龄是:", name)
	} else {
		fmt.Println("这个人不存在")
	}

	for i, j := range map2 {
		fmt.Println(i, j)
	}
}
