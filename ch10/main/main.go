package main

import (
	"fmt"
	_ "learngo/ch09/user"
	. "learngo/ch10/user"
)

func main() {
	c := Course{
		Name: "go",
	}

	fmt.Println(GetCourse(c))
}
