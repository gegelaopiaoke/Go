package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		strings包下的关于字符串的函数
	*/
	s1 := "hello world"
	//1。是否存在substr中的指定字符
	fmt.Println(strings.Contains(s1, "h"))
	//2.是否包含chars中任意一个字符
	fmt.Println(strings.ContainsAny(s1, "abcd"))
	//3.统计substr中的字符出现了多少次
	fmt.Println(strings.Count(s1, "l"))

	s2 := "20190525课堂笔记.txt"
	//4.以指定xxx前缀开头
	if strings.HasPrefix(s2, "201905") {
		fmt.Println("这是19年5月的文件")
	}
	//5.以指定xxx后缀结尾
	if strings.HasSuffix(s2, ".txt") {
		fmt.Println("这是txt文件")
	}
	//查找substr在s中的位置，不存在返回-1，存在返回第一个的下标
	fmt.Println(strings.Index(s1, "l"))
	//查找chars任意一个字符在s中的位置，如果存在返回下标，全部不存在返回-1
	fmt.Println(strings.IndexAny(s1, "abce"))
	//查找chars字符在s中最后一次出现的位置，如果存在返回下标，全部不存在返回-1
	fmt.Println(strings.LastIndex(s1, "l"))

	//字符串拼接
	ss1 := []string{"abc", "world", "hello"}
	//使用sep字符串拼接到s位部
	s3 := strings.Join(ss1, "*")
	fmt.Println(s3)
	//通过删除sep的字符串来切割s
	s4 := "123,456,789,aaa,49595,45"
	ss2 := strings.Split(s4, ",")
	fmt.Println(ss2)
	//重复，自己拼接自己，次数根据count来
	s5 := strings.Repeat("hello", 5)
	fmt.Println(s5)
	// 替换，把new的值替换掉old的值，n为次数，-1为全部
	s6 := strings.Replace(s1, "l", "*", -1)
	fmt.Println(s6)

	s7 := "heLLo WOrlD***123..."
	fmt.Println(strings.ToLower(s7))
	fmt.Println(strings.ToUpper(s7))
}
