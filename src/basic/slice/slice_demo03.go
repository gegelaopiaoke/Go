package main

import "fmt"

func main() {
	/*
		slice := arr[start:end]
			切片中的的数据：[start,end]
			arr[:end],从头到end
			arr[start:],从start到末尾
	*/
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("-----------------1.已有数组直接创建切片-----------------")
	s1 := a[:5]
	s2 := a[3:8]
	s3 := a[5:]
	s4 := a[:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", s3)
	fmt.Printf("%p\n", s4)

	fmt.Println("-----------------2.长度和容量-----------------")
	fmt.Printf("s1 len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("s1 len:%d,cap:%d\n", len(s2), cap(s2))
	fmt.Printf("s1 len:%d,cap:%d\n", len(s3), cap(s3))
	fmt.Printf("s1 len:%d,cap:%d\n", len(s4), cap(s4))

	fmt.Println("-----------------3.更改数组的内容-----------------")
	a[4] = 100
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println("-----------------4.更改切片的内容-----------------")
	s2[2] = 200
	fmt.Println(a)
	fmt.Println(s2)
	fmt.Println("-----------------5.切片添加的内容-----------------")
	s1 = append(s1, 111, 222, 333, 444, 555)
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

}
