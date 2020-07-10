package main

import "fmt"

func main() {
	/*
		指针：pointer
			存储了另一个变量的内存地址的变量
	*/
	//1.定义一个int类型的变量
	a := 10
	fmt.Println("a的数值是:", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("a的地址是：%p\n", &a)

	//2.创建一个指针变量，用来存储变量a的地址
	var p1 *int
	fmt.Println(p1)           //<nil>空指针
	p1 = &a                   //p1指向了a的地址值
	fmt.Println("p1的数值为", p1) //p1中存储的是a的地址
	fmt.Printf("p1的地址为：%p\n", p1)
	fmt.Println("p1的数值，是a1的地址，该地址的存储数据为：", *p1) //获取指针指向的变量的地址

	//3.操作变量，更改数值
	a = 100
	fmt.Println(a)
	fmt.Printf("a的地址是：%p\n", &a)

	//4.通过指针，改变变量的数值
	*p1 = 200
	fmt.Println(a)
	fmt.Printf("a的地址是：%p\n", &a)

}
