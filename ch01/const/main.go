package main

import "fmt"

func main() {
	const PI = 3.1415926
	//const MY_NAME
	const (
		UNKNOWN = 1
		FEMALE = 2
		MALE = 3
	)

	const (
		x int = 16
		y
		s = "abc"
		z
		m
	)

	fmt.Println(x, y, s, z, m)
}
