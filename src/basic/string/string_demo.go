package main

import "fmt"

func main() {
	/*
		Go中的字符串是一个字节的切片。
			可以通过将其内容封装在""中来创建字符串，Go的字符串是Unicode兼容的，并且是UTF-8编码的
		字符串使一些字节的集合

		语法：""，''

		字符：----->对应编码表中的编码值
		其实就是操作字符的编码值
		A---->65
		B---->66
	*/

	s1 := "hello"
	s2 := `hello world`
	fmt.Println(s1)
	fmt.Println(s2)

	//2.字符串长度
	fmt.Println(len(s1))
	fmt.Println(len(s2))

	//3.获取某个字节
	fmt.Println(s2[0])
}
