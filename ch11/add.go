package ch11

// 单元测试 go test
// go test 命令是一个按照一定约定和组织的测试代码驱动程序
// 在包目录中，所有以_test.go为后缀的源码文件都会被go test 运行到
// 我们写的_test.go源码文件不要担心内容过多，因为go build命令不会将这些测试文件打包到最后的可执行文件中
// test 文件有4类，Test开头的 功能测试Benchmark开头的 性能测试 Example 模糊测试
func add(a, b int) int {
	return a + b
}
