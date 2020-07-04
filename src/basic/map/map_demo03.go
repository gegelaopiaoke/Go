package main

import "fmt"

func main() {
	/*
		map和slice的结合使用：
			1.创建map用于存储人的信息
				name,age,sex,address
			2.每个map存储一个人的信息

			3.将这些map存入到slice中

			4.打印遍历输出
	*/
	map1 := make(map[string]string)
	map1["name"] = "王二狗"
	map1["age"] = "30"
	map1["sex"] = "男"
	map1["address"] = "中国"
	fmt.Println(map1)

	map2 := make(map[string]string)
	map2["name"] = "李小花"
	map2["age"] = "18"
	map2["sex"] = "女"
	map2["address"] = "中国"
	fmt.Println(map2)

	s1 := make([]map[string]string, 0, 2)
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	fmt.Println(s1)

	for i, val := range s1 {
		fmt.Printf("第%d个人的信息是：\n", i+1)
		fmt.Printf("\t姓名:%s\n", val["name"])
		fmt.Printf("\t年龄:%s\n", val["age"])
		fmt.Printf("\t性别:%s\n", val["sex"])
		fmt.Printf("\t地址:%s\n", val["address"])

	}
}
