package main

import (
	"fmt"
	"strconv"
)

func main() {
	// for 循环
	/*
	for init; condition; post {}
	 */
	//var i int
	//for ; i < 3; i++  {
	//	fmt.Println(i)
	//}

	//for i < 3 {
	//	time.Sleep(2 * time.Second)
	//	fmt.Println(i)
	//	i++
	//}

	//var sum int
	//for i := 1; i < 101; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(strconv.Itoa(j) + "*" +strconv.Itoa(i) + "=" + strconv.Itoa(i * j))
			fmt.Print("\t")
		}
		fmt.Println()
	}

}
