package main

import "fmt"

// 我想要放多个person的信息到一个集合中
var persons [][]string


// 个人信息：name、age、address、height
// 类型的集合

type Person struct {
	name string
	age int
	address string
	height float32
}

func main() {
	// 很麻烦
	//persons := append(persons, []string{"bobby", "18", "慕课网", "1.80"})

	// 如何初始化结构体
	//p1 := Person{
	//	"bobby1",
	//	18,
	//	"慕课网",
	//	1.80
	//}
	//
	//p2 := Person {
	//	name: "bobby2",
	//	age: 18,
	//	address: "北京市",
	//	height: 1.78,
	//}
	//
	//var persons []Person
	//persons = append(persons, p1)
	//persons = append(persons, Person {
	//	name: "bobby3",
	//})
	//
	//persons2 := []Person{
	//	{
	//		name: "b4",
	//	},
	//	{
	//		age: 19,
	//	},
	//}

	var p Person
	p.age = 20

	fmt.Println(p.age)
	fmt.Println(p.height)

	// 匿名结构体，匿名函数
	address := struct {
		province string
		city string
		address string
	}{
		"北京市",
		"通州区",
		"xxx",
	}

	fmt.Println(address.province)
}
