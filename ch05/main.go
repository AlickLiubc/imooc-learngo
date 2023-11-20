package main

import "fmt"

func main() {
	// go语言提供了哪些集合类型的数据结构，
	// 数组，切片(slice)，map，list
	// 数组 定义：var name [count]int

	// 只有3个元素的数组
	// var course1 [3]string

	// []string [3]string 是2种不同的类型
	// 只有4个元素的数组
	// var course2 [4]string

	//course1[0] = "go"
	//course1[1] = "grpc"
	//course1[2] = "gin"

	// fmt.Println(course1)
	// fmt.Printf("%T\r\n", course1)
	// fmt.Printf("%T", course2)

	//for _, value := range course1 {
	//	fmt.Println(value)
	//}

	// 数组的初始化
	// course1 := [3]string{"go", "grpc", "gin"}


	//course1 := [3]string{2: "grpc"}
	//for _, value := range course1 {
	//	fmt.Println(value)
	//}

	course1 := [...]string{"go", "grpc", "gin"}
	for _, value := range course1 {
		fmt.Println(value)
	}

	for i := 0; i < len(course1); i++ {
		fmt.Println(course1[i])
	}

	course2 := [...]string{"go", "grpc"}
	course3 := [...]string{"go", "gin"}

	if course2 == course3 {
		fmt.Println("equal")
	}

	// 多维数组
	var courseInfo [3][4]string
	courseInfo[0] = [4]string{"go", "1h", "bobby", "go体系课"}
	courseInfo[1] = [4]string{"grpc", "2h", "bobby2", "grpc入门"}
	courseInfo[2] = [4]string{"gin", "1.5h", "bobby3", "gin高级开发"}

	for i := 0; i < len(courseInfo); i++ {
		for j := 0; j < len(courseInfo[i]); j++ {
			fmt.Print(courseInfo[i][j] + "\t")
		}

		fmt.Println()
	}

	for _, row := range courseInfo {
		fmt.Println(row)
		//for _, column := range row {
		//	fmt.Print(column + "\t")
		//}
		//
		//fmt.Println()
	}
}
