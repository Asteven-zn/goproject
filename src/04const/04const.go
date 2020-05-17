package main

import "fmt"

func main() {
	//const定义一个常量，常量的值不可以改变
	const a = 5

	//批量声明常量
	const (
		statusOK = 200
		notFoud  = 404
	)

	//批量声明常量,如果后面没有些值，和上面的常量值相同
	const (
		n1 = 100
		n2
		n3
	)
	fmt.Printf("a type is %T,a = %d\n", a, a)
	fmt.Printf("statusOK is %d,notFoud is %d\n", statusOK, notFoud)
	fmt.Printf("n1 = %d,n2 = %d,n3 = %d\n", n1, n2, n3)
}
