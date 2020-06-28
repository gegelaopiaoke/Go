package main

import "fmt"

func main() {
	/*
		变量：variable
		概念：一小块内存，用于存储数据，在程序运行过程中数值可以改变

		使用：
			step1:变量的声明，也叫定义
			step2:变量的访问，赋值和取值

		go的特性
			静态语言：强类型语言
				go，java，C++，C#
			动态语言：弱类型语言
				JavaScript，php，python，ruby
	*/
	// 第一种： 定义变量，然后赋值
	var num1 int
	num1 = 30
	fmt.Printf("Num1的数值为%d\n", num1)
	var num2 int = 15
	fmt.Printf("num2的数值是%d\n", num2)
}
