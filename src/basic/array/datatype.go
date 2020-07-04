package main

import "fmt"

func main() {
	/*
		Go语言基本数据类型：
		1.基本数据类型：
			布尔类型:bool
				取值：true,false
			数值类型:
				整数：
				浮点：小数
				复数：complex
			字符串：string
		2.复合数据类型
			array,slice,map,function,pointer,struct,interface,channel...
	*/
	//1.布尔类型
	var b1 bool
	b1 = true
	fmt.Printf("%T,%t\n", b1, b1)
	b2 := false
	fmt.Printf("%T,%t\n", b2, b2)

}
