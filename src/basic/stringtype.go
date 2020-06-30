package main

import "fmt"

func main() {
	/*
		字符串：
		1.概念：多个byte的集合，理解为一个字符序列
		2.语法：使用双引号
			"abc","hello","A"
				也可以使用
	*/
	//1.定义字符串
	var s1 string
	s1 = "王二狗"
	fmt.Printf("%T,%s\n", s1, s1)

	s2 := `Hello world`
	fmt.Printf("%T,%s\n", s2, s2)
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T\n", v1)
	fmt.Printf("%T,%s\n", v2, v2)

}
