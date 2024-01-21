package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{
		name: "bobby",
		age:  18,
	}

	p2 := Person{
		name: "bobby",
		age:  18,
	}

	if p1 == p2 {
		fmt.Println("yes")
	}

	//var ps []Person
	var ps2 = make([]Person, 0)
	if ps2 == nil {
		fmt.Println("nil slice")
	}

	var m1 map[string]string
	var m2 map[string]string = make(map[string]string, 0)
	//if m2 == nil {
	//	fmt.Println("nil map")
	//}

	for key, value := range m1 {
		fmt.Println(key, value)
	}

	fmt.Println(m1["name"])
	// m1["name"] = "bobby1"
	m2["name"] = "bobby2"

	for key, value := range m2 {
		fmt.Println(key, value)
	}

}
