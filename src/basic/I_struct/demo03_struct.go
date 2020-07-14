package main

import "fmt"

func main() {
	/*
		匿名结构体和匿名字段

		匿名结构体：没有名字的结构体
			在创建匿名结构体的同时创建对象
			变量名:=struct{
				定义字段Field
			}{
				字段进行赋值
			}
		匿名字段：一个结构体的字段没有字段名

		匿名函数：

	*/
	s1 := Student{"张三", 18}
	fmt.Println(s1)

	func() {
		fmt.Println("hello world")
	}()

	s2 := struct {
		name string
		age  int
	}{
		"李四",
		19,
	}

	fmt.Println(s2)
	//
	//w1 := Worker{"王二狗", 30}
	//fmt.Println(w1)
}

type Worker struct {
	name string
	age  int
}

type Student struct {
	name string
	age  int
}
