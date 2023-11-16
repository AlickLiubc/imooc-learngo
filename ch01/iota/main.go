package main

import "fmt"

func a()(int, bool)  {
	return 0, false
}

// 全局
var name = "bobby"

func main() {
	// 匿名变量
	var _ int
	_, ok := a()

	if ok {
		// 打印
	}

	var mainName = "main"
	fmt.Println(mainName)

	const (
		ERR1 = iota
		//ERR2 = iota
		//ERR3 = iota
		//ERR4 = iota
		ERR2
		ERR3
		ERR4
	)

	fmt.Println(ERR1, ERR2, ERR3, ERR4)

	{
		localName := "local"
		fmt.Println(localName)
	}

	var mname string
	if name == "bobby" {
		mname = "imooc"
	}

	fmt.Println(mname)
}
