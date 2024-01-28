package main

import (
	"fmt"
	"learngo/ch10/user"
)

func main() {
	c := user.Course{
		Name: "go",
	}

	fmt.Println(user.GetCourse(c))
}
