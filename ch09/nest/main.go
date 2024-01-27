package main

import "fmt"

type MyReader interface {
	read()
}

type MyWriter interface {
	write(string)
}

type MyReaderWriter interface {
	MyReader
	MyWriter
	readWrite()
}

type SreaderWriter struct{}

func (s SreaderWriter) read() {
	fmt.Println("read")
}

func (s SreaderWriter) write(s2 string) {
	fmt.Println("write")
}

func (s SreaderWriter) readWrite() {
	fmt.Println("read and write")
}

func main() {
	var srw MyReaderWriter = &SreaderWriter{}
	srw.read()
}
