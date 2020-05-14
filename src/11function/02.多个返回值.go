package main

import "fmt"

//多个返回值
func myfunc01() (int, int, int) {
	return 1, 2, 3
}

//go官方推荐方法
func myfunc02() (e int, f int, g int) {
	e, f, g = 111, 222, 333
	return
}

func main () {
	//函数调用
	a, b, c := myfunc01()
	fmt.Printf("a = %d, b = %d, c = %d\n",a, b, c)

	e, f, g := myfunc02()
	fmt.Printf("e = %d, f = %d, g = %d\n",e, f, g)
}