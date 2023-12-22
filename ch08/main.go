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

	// var a int = 10
	// b := &a

	// 指针需要初始化
	// var b *int
	// var p *Person // 不行 指针需要初始化
	// fmt.Println(p.name)
	// var ps Person

	// 第1种初始化方式
	// ps := &Person{}

	// 第2种初始化方式
	// var emptyPerson Person
	// pi := &emptyPerson

	// 初始化2个关键字,map,channel,slice初始化使用make
	// 指针初始化推荐使用new函数，指针要初始化，否则会出现nil
	// map必须初始化

	// 第3种初始化方式
	var pp = new(Person)
	fmt.Println(pp.name)

}
