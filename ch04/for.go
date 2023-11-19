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
			fmt.Print(strconv.Itoa(j) + "*" + strconv.Itoa(i) + "=" + strconv.Itoa(i*j))
			fmt.Print("\t")
		}
		fmt.Println()
	}

	// for 循环还有一种用法，for range，主要就是对字符串、数组、切片等
	/*
		for key, value := range {

		}
	*/
	name := "imooc go"
	for key, value := range name {
		// fmt.Println(key, value)
		fmt.Printf("%d, %c\r\n", key, value)
	}

	/*
		字符串：key代表字符串索引,value是值的拷贝
		数组
		切片
		map
		channel
	*/
	// for range key, value

}
