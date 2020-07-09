package main

import "fmt"

func main() {
	/*
			defer的词义："推迟"，"延迟"
			在go语言中，使用defer关键字来延迟一个函数或者方法的执行
		1.在defer函数或者方法：一个函数或者方法的执行被延迟了
		2.defer的用法：
			A：对象.close(),临时文件的删除
			B：go语言中关于异常的处理，使用panic()和recover()
				panic函数用于引发恐慌，导致程序中断执行
				recover函数用于恢复程序的执行，recover()语法要求必须在defer中执行
		3.如果有多个defer函数,先延迟的后执行，后延迟的限制性。
		4.defer函数传递参数的时候：defer函数调用时，就已经传递了参数数据，只是暂时不执行函数的代码而已。
		5.defer函数的注意点：
			当外围函数的语句正常执行完毕时，只是其中所有的延迟函数都执行完毕后，外围函数才会真正的结束执行
			当执行外围函数的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正的返回。
			当外围函数中的代码引发恐慌的时候，只有其中所有的延迟函数都执行完毕，该运行时的恐慌才会真正被扩展至调用函数

	*/

	//defer fun1a("hello")
	//defer fmt.Println("12345")
	//defer fun1a("world")
	//defer fmt.Println("王二狗")

	a := 2
	fmt.Println(a)
	defer fun2a(a)
	a++
	fmt.Println(a)

}

func fun2a(a int) {
	fmt.Println(a)

}
func fun1a(s string) {
	fmt.Println(s)
}
