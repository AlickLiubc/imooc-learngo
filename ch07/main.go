package main

import (
	"fmt"
	"strconv"
)

// 自定义类型，基于已有的类型自定义一个类型
type MyInt int

func (mi MyInt) string() string {
	return strconv.Itoa(int(mi))
}


func main() {
	// type关键字
	/**
	1. 定义结构体
	2. 定义接口
	3. 定义类型别名
	4. 类型定义
	5. 类型判断
	 */
	//// 别名实际上，是为了更好地理解代码
	//type MyInt = int	// 类型别名
	//var i MyInt = 12
	//var j int = 8
	//// 在编译的时候，类型别名会被直接替换为int
	//fmt.Println(i + j)
	//fmt.Printf("%T\r\n", i)

	//type MyInt int	// 自定义类型,基于已有的类型自定义类型

	var a interface{} = "abc"
	switch a.(type) {
		case string:
			fmt.Println("字符串")
	}

	var i MyInt = 12
	fmt.Println(i.string())

	// var j int = 8
	// 既希望你是int类型，又希望增加方法
	// fmt.Println(int(i) + j)
	fmt.Printf("%T\r\n", i)
}
