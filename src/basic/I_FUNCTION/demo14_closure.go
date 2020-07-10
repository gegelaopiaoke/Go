package main

import "fmt"

func main() {
	/*
			go语言支持函数式编程
				支持将一个函数作为另一个函数的参数，
				也支持将一个函数作为另一个函数的返回值

		闭包(closure)：
			一个外层函数中，有内层函数，该内层函数中，会操作外部函数的局部变量(外层函数中的参数，或者外层函数中直接定义的变量)
			并且该外层函数的返回值就是这个内层函数
			这个内层函数和外部函数的局部变量，统称为闭包结构
	*/
	res1 := increment()      //res1=fun
	fmt.Printf("%T\n", res1) //类型为func() int
	fmt.Println(res1)
	v1 := res1()
	fmt.Println(v1)
	v2 := res1()
	fmt.Println(v2)
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())

	res2 := increment()
	fmt.Println(res2)
	fmt.Println(res2())
	fmt.Println(res1())

}
func increment() func() int { //外层函数
	//1.定义了一个局部变量
	i := 0
	//2.定义了一个匿名函数，给变量自增返回
	fun := func() int {
		i++
		return i
	}
	//3.返回该匿名函数
	return fun
}
