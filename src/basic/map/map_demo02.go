package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]string)
	map1[1] = "A"
	map1[2] = "B"
	map1[3] = "C"
	map1[4] = "D"
	map1[5] = "E"

	for k, v := range map1 {
		fmt.Println(k, v)
	}
	fmt.Println("-------------")

	for i := 1; i <= len(map1); i++ {
		fmt.Println(i, "----->", map1[i])
	}
	/*
		1.获取所有的key，--->切片/数组
		2.进行排序
		3.遍历key，---->map[key]
	*/
	keys := make([]int, 0, len(map1))
	fmt.Println(keys)
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println("-------------")
	fmt.Println(keys)

	//冒泡排序，或者使用sort包下的排序方法
	sort.Ints(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, map1[key])
	}
	fmt.Println("-------------")
	s1 := []string{"Apple", "Windows", "Orange", "abc", "王二狗", "abc", "aaa"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)
}
