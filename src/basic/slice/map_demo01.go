package main

import "fmt"

func main() {
	/*
			map: 映射，是一种专门用于存储键值对的集合，属于引用类型

			存储特点：
				A：存储的是无序的键值对
				B：键不呢给你充个农夫，并且和value值一一对应。
					map中的key不能重复，如果重复，那么新的value会覆盖原来的，程序不会报错

		语法结果：
			1.创建map

		每种数据类型的默认值：
			int：0
			float：0.0 -->简写成0
			string:""
			array:[]    存储你这种类型的默认值

			slice：nil  因为有底层数组，所以可以直接使用
			map：nil    也就是null，空的map不可以直接使用

	*/

	//创建map
	var map1 map[int]string
	var map2 = make(map[int]string)
	var map3 = map[string]int{"Go": 98, "Python": 87, "Java": 79, "Html": 93}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

}
