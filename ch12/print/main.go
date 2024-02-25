package main

import (
	"fmt"
	"time"
)

var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1

	for {
		// 我怎么去做到，应该此处，等待另一个goroutine来通知我
		<-number
		fmt.Printf("%d%d", i, i+1)
		i += 2
		letter <- true
	}
}

func printLetter() {
	i := 0

	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for {
		<-letter
		if i >= len(str) {
			return
		}
		fmt.Printf(str[i : i+2])
		i += 2
		number <- true
	}
}

func main() {

	go printNum()
	go printLetter()

	number <- true

	time.Sleep(10 * time.Second)
}
