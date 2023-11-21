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

	// 切片的初始化 3种
	// 1：从数组直接创建
	// 2：使用string{}
	// 3：make
	allCourses := [5]string{"go", "grpc", "gin", "mysql", "elasticsearch"}
	// 左闭右开的区间
	courseSlice := allCourses[0:len(allCourses)]
	fmt.Println(courseSlice)

	// courseSlice := []string{"go", "grpc", "gin", "mysql", "elasticsearch"}

	// make 函数
	allCourses2 := make([]string, 3)
	allCourses2[0] = "c"
	fmt.Println(allCourses2)

	var allCourse3 []string
	allCourse3[0] = "c"
}
