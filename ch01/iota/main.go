package main

import "fmt"

func a() (int, bool) {
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
		ERR1 = iota + 1
		//ERR2 = iota
		//ERR3 = iota
		//ERR4 = iota
		ERR2
		ERR25 = "ha" // iota内部仍然会增加计数器
		ERR3         // iota+1
		ERR31        // iota+1
		ERR32        // iota+1
		ERR33        // iota+1
		ERR34 = 100
		ERR4  = iota
	)

	const (
		ERRNEW1 = iota
	)

	fmt.Println(ERR1, ERR2, ERR25, ERR3, ERR4)
	fmt.Println(ERRNEW1)

	{
		localName := "local"
		fmt.Println(localName)
	}
	// fmt.Println(localName)

	var mname string
	if name == "bobby" {
		mname = "imooc"
	}

	fmt.Println(mname)
}
