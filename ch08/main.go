package main

import "fmt"

type Person struct {
	name string
}

func changeName(p *Person) {
	p.name = "imooc"
}

func main() {
	// 指针，提一个需求，结构体传值时，函数中修改的值可以反应到变量中
	// p := Person{
	//	 name: "bobby",
	// }

	// var pi *Person = &p

	// changeName(&p)
	// fmt.Println(p.name)

	// fmt.Printf("%p", pi)

	// 指针的定义
	// var po *Person
	// po = &p
	// fmt.Println(po)

	po := &Person{
		name: "bobby2",
	}

	(*po).name = "bobby3"
	po.name = "bobby4"

	// 第一个不同的点就出来，第二个点语言限制了指针的运算
	// go的指针不同于C语言的指针运算，unsafe包里面，所以，一般我们不会使用unsafe包
	fmt.Println(po)
}
