package main

import "fmt"

func main() {
	/*
		return语句：词义"返回"
			A：一个函数有返回值，那么使用return将返回值给调用处
			B：同时意味着结束了函数的运行
		注意点：
			1.一个函数定义了返回值，必须使用return语句将返回结果给调用处，return后的数据必须和函数定义的数据一致：个数，类型，顺序
			2.可以使用_,来舍去多余的返回值
			3.如果一个函数定义来有返回值，那么函数中有分支，循环，那么要保证，无论执行来哪个分支，都要有return语句被执行到
	*/
	a, b, c := fun111()
	fmt.Println(a, b, c)
	_, _, d := fun111()
	fmt.Println(d)
	fmt.Println(fun222(-50))
	fun333()
}
func fun111() (int, float64, string) {
	return 0, 0, "aa"
}

func fun222(age int) int {
	if age >= 0 {
		return age
	} else {
		fmt.Println("不能为负数")
		return 0
	}
}

func fun333() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			return
		}
		fmt.Println(i)
	}

	fmt.Println("fun3()...over...")
}
