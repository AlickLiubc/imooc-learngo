package main

import (
	"fmt"
	"strconv"
)

func printSlice(data []string) {
	data[0] = "java"

	for i := 0; i < 10; i++ {
		data = append(data, strconv.Itoa(i))
	}
}

func main()  {
	// go的slice在函数参数传递的时候是值传递，还是引用传递
	// 值传递，效果又呈现了引用传递
	courses := []string{"go", "grpc", "gin"}
	printSlice(courses)

	fmt.Println(courses)
}
