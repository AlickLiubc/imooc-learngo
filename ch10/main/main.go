package main

import (
	"fmt"
	_ "learngo/ch09/user"
	. "learngo/ch10/user" // 导入的别名 这种用法尽量少用
)

func main() {
	c := Course{
		Name: "go",
	}

	fmt.Println(GetCourse(c))
}
