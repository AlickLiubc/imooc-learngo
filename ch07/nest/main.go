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

// 接收器有2种形态
//func (p Person) print() {
func (p *Person) print() {
	// 有可能该函数中想要改变p的值，就是person对象很大，
	// 数据较大的时候，考虑使用指针方式传递
	p.age = 20
	fmt.Printf("name: %s, age: %d\r\n", p.name, p.age)
}

// func (s sturctType) funcName(param1 param1Type,...)(returnVal1 returnType1,...){}

func main() {
	p := Person{
		"bobby", 18,
	}
	p.print()

	fmt.Println(p.age)

	//s := Student {
	//	Person {
	//		"bobby", 19,
	//	},
	//	95.6,
	//}
	//
	//s.print()

	// s := Student{}
	// s.p.name = "bobby"
	// fmt.Println(s.p.age)
	// fmt.Println(s.age)
}

