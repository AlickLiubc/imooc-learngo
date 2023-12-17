package main

import "fmt"

type Person struct {
	name string
	age int
}

type Student struct {
	// 第1种嵌套方式
	// p Person

	// 第2种嵌套方式，匿名嵌套
	Person
	score float32
}

func main() {
	//s := Student {
	//	Person {
	//		"bobby", 18,
	//	},
	//	95.6,
	//}

	s := Student{}
	// s.p.name = "bobby"
	// fmt.Println(s.p.age)
	fmt.Println(s.age)
}

