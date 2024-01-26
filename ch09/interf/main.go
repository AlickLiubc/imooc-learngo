package main

import "fmt"

type Duck interface{
	// 方法的声明
	Gaga()
	Walk()
	Swim()
}

type PskDuck struct {
	legs int
}

func (pd *PskDuck) Gaga() {
	fmt.Println("This PskDuck Gaga")
}

func (pd *PskDuck) Walk() {
	fmt.Println("This PskDuck Walk")
}

func (pd *PskDuck) Swim() {
	fmt.Println("This PskDuck Swim")
}

func main()  {
	var duck Duck = &PskDuck{}
	duck.Gaga()
}
