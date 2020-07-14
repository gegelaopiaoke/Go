package main

import "fmt"

func main() {
	/*
		数据类型：
			值类型：int,float,bool,string,array,struct

			引用类型：slice,map,function,pointer

		通过指针
			new()

	*/
	//1.结构体是值类型
	p1 := Person{"王二狗", 30, "男", "北京市"}
	fmt.Println(p1)
	fmt.Printf("%p,%T\n", &p1, p1)
	p2 := p1
	fmt.Println(p2)
	fmt.Printf("%p,%T\n", &p2, p2)
	p2.name = "李小花"
	fmt.Println(p2)
	fmt.Println(p1)
	fmt.Println("---------------------------------")
	//2.定义结构体指针
	var pp1 *Person
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Printf("%p,%T\n", pp1, pp1)
	fmt.Println(*pp1)
	pp1.name = "李四"
	fmt.Println(*pp1)
	fmt.Println(p1)
	fmt.Println("---------------------------------")

	//3.使用内置函数new()，go语言中专门用于创建某种类型的指针函数
	pp2 := new(Person)
	fmt.Println(pp2)
	pp2.name = "王五"
	fmt.Println(pp2)
	fmt.Printf("%p,%T\n", pp2, pp2)
	fmt.Println(*pp2)
	fmt.Println("---------------------------------")
	pp3 := new(int)
	fmt.Println(pp3)
	fmt.Printf("%p,%T\n", pp3, pp3)
}

type Person struct {
	name    string
	age     int
	sex     string
	address string
}
