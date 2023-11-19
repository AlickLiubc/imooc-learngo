package main

import "fmt"

func main() {
	var course1 [3]string
	var course2 [4]string

	fmt.Printf("%T\r\n", course1)
	fmt.Printf("%T", course2)
}
