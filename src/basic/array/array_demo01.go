package main

import "fmt"

func main() {
	//step1: 创建数组
	var arr1 [4]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	fmt.Println(arr1[0], arr1[1])
	fmt.Println("数组的长度为：", len(arr1))
	fmt.Println("数组的容量为：", cap(arr1))
	arr1[0] = 100
	fmt.Println(arr1[0], arr1[1])
	var a [4]int
	fmt.Println(a)
	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b)
	var c = [5]int{1, 2, 3}
	fmt.Println(c)
	//下标为1的地方存储1，下标为3的地方存储2
	var d = [5]int{1: 1, 3: 2}
	fmt.Println(d)

	var num1 int
	num1 = 100
	fmt.Printf("%p\n", &num1)
}
