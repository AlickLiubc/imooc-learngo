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
	//courses := []string{"go", "grpc", "gin"}
	//printSlice(courses)
	//
	//fmt.Println(courses)
	//
	//courses = append(courses, "abc")

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9 ,10}
	s1 := data[1:6]
	s2 := data[2:7]

	s2 = append(s2, 1, 2, 3, 4, 5 ,6 ,7, 8, 9, 1, 1, 1, 1)
	s2[0] = 22

	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))

	fmt.Println(s2)
	fmt.Println(s1)

	var dataInt []int
	for i := 1; i < 2000; i++ {
		dataInt = append(dataInt, i)

		fmt.Printf("len:%d, cap:%d\r\n", len(dataInt), cap(dataInt))
	}


}
