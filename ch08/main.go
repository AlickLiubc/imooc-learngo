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
	p := Person {
		name: "bobby",
	}

	changeName(&p)

	fmt.Println(p.name)
}
