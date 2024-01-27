package main

import (
	"fmt"
	"strings"
)

func add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi
	case float64:
		ai, _ := a.(float64)
		bi, _ := b.(float64)
		return ai + bi
	case string:
		as, _ := a.(string)
		bs, _ := b.(string)
		return as + bs
	default:
		panic("不支持的类型")
	}
}

func main() {
	a := "hello "
	b := "world"

	re := add(a, b)
	res, _ := re.(string)

	fmt.Println(re)
	fmt.Println(strings.Split(res, " "))
}
