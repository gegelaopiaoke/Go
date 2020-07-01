package main

import "fmt"

func main() {
	num := 10
	fmt.Printf("%T\n", num)

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [3]float64{2.15, 3.18, 6.19}
	arr3 := [4]int{5, 6, 7, 8}
	arr4 := [2]string{"hello", "world"}
	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)
	fmt.Printf("%T\n", arr3)
	fmt.Printf("%T\n", arr4)

	num2 := num
	fmt.Println(num, num2)
	num2 = 20
	fmt.Println(num, num2)

}
