package main

import "fmt"

func main() {
	//// map是一个key(索引)和value(值)的无序集合，主要是查询方便O(1)
	//var courseMap = map[string]string{
	//	"go": "go工程师",
	//	"grpc": "grpc入门",
	//	"gin": "gin深入理解",
	//}
	//
	//// 取值
	//fmt.Println(courseMap["grpc"])
	//
	//// 放值
	//courseMap["mysql"] = "mysql的原理"
	//
	//fmt.Println(courseMap["mysql"])

	// nil，map类型想要设置值必须要先初始化
	//var courseMap map[string]string{}
	var courseMap = make(map[string]string, 3)	// make是内置函数，主要用于初始化slice, map, channel
	courseMap["mysql"] = "mysql的原理"
	fmt.Println(courseMap)
	// panic: assignment to entry in nil map

	// map必须初始化才能进行使用
	// 1.map[string]string{}
	// 2.make(map[string]string, 3)
	// 3.但是slice可以不初始化

	//var m []string
	//if m == nil {
	//	fmt.Println("yes")
	//}
	//m = append(m, "a")
	//fmt.Println(m)

	// 遍历
	//for key,value := range courseMap {
	//	fmt.Println(key, value)
	//}

	//for _, value := range courseMap {
	//	fmt.Println(value)
	//}

	for key := range courseMap {
		fmt.Println(key, courseMap[key])
	}

}
