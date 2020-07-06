package main

import "fmt"

func main() {
	/*
		可变参数：
			概念：一个函数的参数的类型确定，但是个数不确定，就可以使用可变参数
		语法：
			参数名	... 参数类型

			对于函数，可变参数相当于一个切片
			调用那个函数的时候，可以传入0-多个函数

			Println(),Printf(),Print()
			append()

		注意事项:
			A:如果一个函数的参数是可变参数，同时还有其他的参数，可变参数要放在参数列表的最后
			B:一个函数的参数列表中最多只能有一个可变参数
	*/
	getSum(1, 2, 3, 4, 5)
	//2.切片
	s1 := []int{1, 2, 3, 4, 5}
	getSum(s1...)
}
func getSum(nums ...int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println(sum)
}
