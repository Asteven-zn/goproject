package main

import "fmt"

//无参有返回值：只有一个返回值
//有返回值的函数需要通过return中断函数，通过return返回
func myfunc01() int {
	return 666
}

//给返回值起一个变量名，go的推荐写法
//常用写法
func myfunc02() (result int) {
	result = 777
	return
}

func main () {
	//无参数有返回值函数的调用
	var a int
	a = myfunc01()
	fmt.Println("a = ", a)

	b := myfunc01()
	fmt.Println("b = ", b)

	c := myfunc02()
	fmt.Println("c = ", c)
}