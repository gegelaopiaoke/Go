package main

import "fmt"

func main() {
	/*
		结构体：是由一系列具有相同类型或者不同类型的数据结构的数据集合
			结构体成员是由一系列的成员变量构成，这些成员变量也被称之为字段
	*/
	var p1 Person
	fmt.Println(p1)
	p1.name = "王二狗"
	p1.age = 30
	p1.sex = "男"
	p1.address = "北京市"
	fmt.Printf("姓名：%s,年龄:%d,性别：%s,地址：%s\n", p1.name, p1.age, p1.sex, p1.address)
	p2 := Person{}
	p2.name = "Ruby"
	p2.age = 28
	p2.sex = "女"
	p2.address = "上海市"
	fmt.Printf("姓名：%s,年龄:%d,性别：%s,地址：%s\n", p2.name, p2.age, p2.sex, p2.address)
}

//1.定义结构体
type Person struct {
	name    string
	age     int
	sex     string
	address string
}
