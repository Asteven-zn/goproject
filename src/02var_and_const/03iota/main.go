package main

import "fmt"

func main() {
	//iota是go语言的常量计数器，只能在常量的表达式中使用
	//在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次
	const (
		a1 = iota //0
		a2        //1
		a3        //2
	)
	fmt.Println(a1, a2, a3)

	//过渡
	const (
		b1 = iota //0
		b2        //1
		_         //2
		b3        //3
	)
	fmt.Println(b1, b2, b3)

	//插队
	const (
		c1 = iota //0
		c2 = 100  //100
		c3        //100
		c4        //100
	)
	fmt.Println(c1, c2, c3, c4)

	const (
		d1 = iota //0
		d2 = 100  //100
		d3 = iota //2
		d4        //3
	)
	fmt.Println(d1, d2, d3, d4)

	const (
		e1, e2 = iota + 1, iota + 2 //e1:1   e2:2
		e3, e4 = iota + 1, iota + 2 //e3:2   e4:3
	)
	fmt.Println(e1, e2, e3, e4)
}
