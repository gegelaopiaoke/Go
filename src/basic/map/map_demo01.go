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
				var map1 map[kay类型]value类型
					nil map，无法直接使用

				var map2 = make(map[key类型])value类型
				注意map中的key不能重复，如果重复，那么新的value会覆盖原来的，程序不会报错
				var map3 = map[key类型]value类型{key:value,key:value,key:value}
			2.修改/添加
				map[key]=value
					如果key不存在，就是添加数据
					如果key存在，就是修改数据

			3.获取
				map[key]-->value

				value,ok:=map[key]
					根据key获取对应的value，这样写的好处是可以判断是0还是不存在的零值



		每种数据类型的默认值：
			int：0
			float：0.0 -->简写成0
			string:""
			array:[]    存储你这种类型的默认值

			slice：nil  因为有底层数组，所以可以直接使用
			map：nil    也就是null，空的map不可以直接使用

	*/

	//1.创建map
	var map1 map[int]string
	var map2 = make(map[int]string)
	var map3 = map[string]int{"Go": 98, "Python": 87, "Java": 79, "Html": 93}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)
	//2.nil map
	if map1 == nil {
		map1 = make(map[int]string)
		fmt.Println(map1 == nil)
	}

	//3.存储键值对到map中
	map1[1] = "hello"
	map1[2] = "world"
	map1[3] = "memeda"
	map1[4] = "王二狗"

	//4.获取数据，根据key获取对应的value
	//如果key存在，获取对应的数值，如果不存在，取0
	fmt.Println(map1)
	fmt.Println(map1[4]) //根据key获取对应的value值

	v1, ok := map1[40]
	if ok {
		fmt.Println("对应的数值是：", v1)
	} else {
		fmt.Println("操作的key不存在，获取的是零值", v1)
	}
	//5.修改数据
	fmt.Println(map1)
	map1[3] = "如花"
	fmt.Println(map1)

	//6.删除数据
	delete(map1, 3)

}
