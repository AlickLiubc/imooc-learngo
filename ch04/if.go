package main

import "fmt"

func main() {
	// if条件判断
	age := 18
	country := "USA"
	if age < 18 {
		if country == "中国" {
			fmt.Println("未成年")
		}
	} else if age == 18 {
		fmt.Println("恰好成年")
	} else {
		fmt.Println("已成年")
	}


	if age < 18 {
		fmt.Println("未成年")
	}

	if age == 18 {
		fmt.Println("恰好成年")
	}

	if age > 18 {
		fmt.Println("已成年")
	}

}
