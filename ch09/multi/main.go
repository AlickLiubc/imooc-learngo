package main

import "fmt"

type MyWriter interface {
	Write(string) error
}

type MyCloser interface {
	Close() error
}

type WriterCloser struct {
	MyWriter	// interface也是一个类型
}

type FileWriter struct {
	filePath string
}

type DataBaseWriter struct {
	username string
	password string
	port string
}

type Writer struct {}

func (fw *FileWriter) Write(string) error {
	fmt.Println("write string to file")
	return nil
}

func (dw *DataBaseWriter) Write(string) error {
	fmt.Println("write string to database")
	return nil
}

//func (wc *WriterCloser) Write(string) error {
//	fmt.Println("write string")
//	return nil
//}

func (wc *WriterCloser) Close() error {
	fmt.Println("close")
	return nil
}

func main() {
	//var mw MyWriter = &WriterCloser{}
	//var mc MyCloser = &WriterCloser{}
	var mw MyWriter = &WriterCloser{
		&FileWriter{},
	}

	mw.Write("bobby")
}
