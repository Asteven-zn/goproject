package main

import (
	"fmt"
	"strings"
)

//累加器
func AddUpper() func (int) int {
	var n int = 10
	var str = "hello"
	return func (x int) int {
		n = n + x
		str += string(36)  //asci => 36 = $
		fmt.Println("str=", str)
		return n
	}
}

//判断文件的后缀名
func makeSuffix(suffix string) func (string) string {
	return func (name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}


func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))

	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f2("weinter"))
	fmt.Println("文件名处理后=", f2("bird.jpg"))
}