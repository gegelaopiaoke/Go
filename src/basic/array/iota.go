package main

import "fmt"

func main() {
	/*
		iota:特殊的常量，可以被编辑器自动修改的常量
			每当定义一个const，iota的初始值为0
			每当定义一个常量，就会自动累计加1
			直到下一个const出现，清零
	*/
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	const (
		d = iota
		e
	)
	fmt.Println(d, e)
	const (
		MALE = iota
		FEMALE
		UNKNOW
	)
	fmt.Println(MALE, FEMALE, UNKNOW)
}
