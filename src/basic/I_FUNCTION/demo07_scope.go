package main

import "fmt"

func main() {
	/*
		作用域：变量可以使用的范围
			局部变量：函数内部定义的变量，就叫局部变量

			全局变量：函数外部定义的变量，就叫做全局变量
	*/
	//定义在main函数中，所以n的作用域就是在main函数的范围内
	n := 10
	fmt.Println(n)
	a := 1
	fmt.Println(a)
	fmt.Println(n)
}
