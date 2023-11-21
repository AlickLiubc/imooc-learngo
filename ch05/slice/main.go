package main

import "fmt"

func main()  {
	// go折中，slice切片 - 数组
	var courses []string
	fmt.Printf("%T\r\n", courses)

	// 这个方法很特别，原理
	courses = append(courses, "go")
	courses = append(courses, "grpc")
	courses = append(courses, "gin")
	fmt.Println(courses)

	fmt.Println(courses[1])

}
