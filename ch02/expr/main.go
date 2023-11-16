package main

import "fmt"

func main()  {
	var a, b = 1, 3
	fmt.Println(a + b)

	var astr, bstr = "hello", "bobby"
	fmt.Println(astr + bstr)

	fmt.Println(3 % 2)
	fmt.Println(4 % 2)
	a++
	fmt.Println(a)

	// 逻辑运算符
	var abool, bbool = true, false
	if abool && bbool {

	}

	if abool || bbool {

	}

	if !abool {

	}

	// 位运算符, 性能要求高的时候一般会考虑与运算
	var A = 60
	var B = 13
	fmt.Println(A & B)
}
