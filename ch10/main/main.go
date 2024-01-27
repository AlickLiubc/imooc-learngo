package main

import (
	"fmt"
	"learngo/ch10/user"
)

func main() {
	c := course.Course{
		Name: "go",
	}

	fmt.Println(course.GetCourse(c))
}
