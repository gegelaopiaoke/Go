package main

import "fmt"

func main() {
	a2 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(a2)
	fmt.Printf("二维数组的地址:%p\n", &a2)
	fmt.Println(a2[0][3])

	for i := 0; i < len(a2); i++ {
		for j := 0; j < len(a2[i]); j++ {
			fmt.Println(a2[i][j], "\t")
		}
	}
	for _, arr := range a2 {
		for _, val := range arr {
			fmt.Println(val, "\t")
		}
	}
}
