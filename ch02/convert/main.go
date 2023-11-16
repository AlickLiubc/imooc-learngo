package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var a int8 = 12
	//var b = uint8(a)
	//
	//var f float32 = 3.14
	//var c = int32(f)

	var f64 = float64(a)
	fmt.Println(f64)

	//type IT int
	//var c IT = IT(a)

	// 字符串转数字
	var istr = "12"
	myint, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(myint)


	var myi = 32
	fmt.Println(strconv.Itoa(myi))

	// 字符串转化为float32,转换为bool
	fmt.Println(strconv.ParseFloat("3.1415", 64))
}
