package main

import "fmt"

func add(a, b interface{}) int {
	ai, _ := a.(int)
	bi, _ := b.(int)

	return ai + bi
}

func main() {
	a := 1
	b := 2
	fmt.Println(add(a, b))
}
