package main

import "fmt"

func main() {
	// 长度计算
	// 如果你想指定一个字符串(中文)长度，如果你只有英文，这个时候直接len
	name := "imooc体系课学习"
	bytes := []rune(name)
	fmt.Println(len(bytes))

	// 转义符
	courseName := `go"体系课"`
	fmt.Println(courseName)

	fmt.Print("hello\r\n")
	fmt.Print("world\r\n")

	// 格式化输出
	username := "bobby"
	fmt.Println("用户名：" + username)
}
