package main

import "fmt"

func main() {
	/*
		Go语言的数据类型：
			数值类型：整数，浮点
				进行运算操作，加减乘除，打印
			字符串：
				可以获取单个字符，截取字符串，遍历，strings包下的函数操作
			数组，切片，map。。
				存储数据，修改数据，获取数据，遍历数据。。。
			函数：
				加()，进行调用
			注意点：
					函数作为一种复合数据类型，可以看作是一种特殊的变量
						函数名():将函数进行调用，函数中的代码会全部执行，然后将return的结果返回调用处
						函数名：指向函数体的地址
	*/
	//1.整形
	a := 10
	a += 5
	fmt.Println("a", a)
	//2.数组，切片，map。。。容器
	b := [4]int{1, 2, 3, 4}
	b[0] = 100
	for i := 0; i < len(b); i++ {
		fmt.Printf("%d\t", b[i])
	}
	fmt.Println()

	//3.函数做一个变量
	fmt.Printf("%T\n", fun1a)
	fmt.Println(fun1a)

	//4.直接定义一个函数类型的变量
	var c func(int, int)
	fmt.Println(c)

	c = fun1a
	fmt.Println(c)
	fun1a(10, 20)
	c(100, 200)
	res1 := fun2
	res2 := fun2(1, 2)
	fmt.Println(res1)
	fmt.Println(res2)
}
func fun2(a, b int) int {
	return a + b

}
func fun1a(a, b int) {
	fmt.Printf("a:%d,b:%d\n", a, b)
}
