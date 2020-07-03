package main

import "fmt"

func main() {
	/*
		切片Slice：
			1.每一个切片引用了一个底层数组
			2.切片本身不存储任何数据，都是这个底层数组存储，所以修改切片也就是修改这个数组中的数据
			3.当向切片添加数据时，如果没有超过容量，直接添加，如果超过容量，自动扩容
	*/
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1) //0xc00000e340

	s1 = append(s1, 4, 5)
	fmt.Println(s1)
	fmt.Printf("%p\n", s1) //0xc00000e340

	s1 = append(s1, 6, 7, 8)
	fmt.Println(s1)
	fmt.Printf("len:%d,cap%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 9, 10)
	fmt.Println(s1)
	fmt.Printf("len:%d,cap%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 11, 12, 13, 14, 15)
	fmt.Println(s1)
	fmt.Printf("len:%d,cap%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25)
	fmt.Println(s1)
	fmt.Printf("len:%d,cap%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

}
