package main

import "fmt"

func main() {
	/*
		函数的返回值：
			一个函数的执行结果，返回给函数的调用处。执行结果就叫函数的返回值
		return语句：
			一个函数的定义上有返回值，那么函数中就必须使用return语句，将结果返回调用处
			函数返回的结果，必须和函数定义的一直：类型，个数，顺序。

		1.将函数的结果返回调用处
		2.同时结束来该函数的执行
	*/
	res := getSum1()
	fmt.Println("返回值为", res)
	fmt.Println(getSum2())

	fmt.Println(rectangle(5, 5))
	fmt.Println(rectangle2(5, 6))
}

//定义一个函数，用来求矩形的周长和面积
func rectangle(len, vid float64) (float64, float64) {
	perimeter := (len + vid) * 2
	area := len * vid
	return perimeter, area
}

func rectangle2(len, vide float64) (peri float64, area float64) {
	peri = (len + vide) * 2
	area = len * vide
	return
}

//定义一个函数，带返回值
func getSum1() int {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	return sum
}

//定义函数时，指明要返回的数据是哪一个
func getSum2() (sum int) {
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return
}
