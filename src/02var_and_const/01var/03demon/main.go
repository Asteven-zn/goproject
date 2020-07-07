package main

import (
	"fmt"
"unsafe"
)

func main() {
	//定义一个变量，并查看变量的数据类型
	var n1 = 100
	fmt.Printf("n1的数据类型是%T\n", n1)

	//定义一个变量，并查看变量的数据类型，并且查看变量占用的字节大小
	//%T值的类型,%d表示为十进制
	var n2 = 9.9
	fmt.Printf("n2的数据类型是%T, n2占用的字节大小是%d\n", n2, unsafe.Sizeof(n2))
}