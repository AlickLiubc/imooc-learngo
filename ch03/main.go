package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	age := 18
	address := "北京"
	mobile := "13800000000"
	// 及其难维护
	fmt.Println("用户名：" + username +
				",年龄：" + strconv.Itoa(age) +
		        ",地址：" + address +
		        ",手机号：" + mobile)

	// 较常用，但是性能没有上一个方式好
	fmt.Printf("用户名：%s,年龄：%d,地址：%s,手机号：%s\r\n", username, age, address, mobile)

	userMsg := fmt.Sprintf("用户名：%T,年龄：%d,地址：%s,手机号：%s\r\n", username, age, address, mobile)
	fmt.Println(userMsg)

	// 通过string的builder进行字符串拼接，性能比较高
	var builder strings.Builder
	builder.WriteString("用户名：")
	builder.WriteString("username")
	builder.WriteString(",年龄：")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString(",地址：")
	builder.WriteString(address)

	re := builder.String()

	fmt.Println(re)
}
