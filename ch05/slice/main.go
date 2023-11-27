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
	// allCourses := [5]string{"go", "grpc", "gin", "mysql", "elasticsearch"}
	//// 左闭右开的区间
	//courseSlice := allCourses[0:len(allCourses)]
	//fmt.Println(courseSlice)
	//
	//// courseSlice := []string{"go", "grpc", "gin", "mysql", "elasticsearch"}
	//
	//// make 函数
	//allCourses2 := make([]string, 3)
	//allCourses2[0] = "c"
	//fmt.Println(allCourses2)
	//
	//var allCourse3 []string
	//allCourse3[0] = "c"

	// 访问切片的元素，可以单个，也可以多个
	// fmt.Println(allCourses[1])

	// [start:end]
	/*
	1. 如果只有start 没有end 就表示从start开始到结尾的所有数据
	2. 如果没有start 有end 就表示从0到end之前的所有数据
	3. 如果有start 没有end
	4. 有start 有end
	*/
	//fmt.Println(allCourses[:])
	//fmt.Println(allCourses[1:])
	//fmt.Println(allCourses[1:3])

	courseSlice := []string{"go", "grpc"}
	courseSlice2 := []string{"gin", "mysql", "elasticsearch"}

	// courseSlice = append(courseSlice, courseSlice2...)
	courseSlice = append(courseSlice, courseSlice2[1:2]...)
	// courseSlice = append(courseSlice, "gin", "mysql", "elasticsearch")
	//for _, value := range courseSlice2 {
	//	courseSlice = append(courseSlice, value)
	//}

	fmt.Println(courseSlice)
}
