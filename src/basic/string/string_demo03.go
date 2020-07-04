package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		strconv包：字符串和基本类型之前的转换
			string convert
	*/
	//fmt.Println("aaa" + 100)
	s1 := "true"
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%t\n", b1, b1)
	ss1 := strconv.FormatBool(b1)
	fmt.Printf("%T,%t\n", ss1, ss1)

}
