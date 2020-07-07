//package main
//
//import "fmt"
//
//func test() bool {
//	fmt.Println("test...")
//	return true
//}
//
//func main() {
//	var i int = 8
//	if i < 9 && test() {
//		fmt.Println("ok...")
//	}
//
//	if i < 9 || test() {
//		fmt.Println("hello...")
//	}
//}
//
////计算乘积
//func calculateBill(price int, num int) int {
//	var totalprice = price * num
//	return totalprice
//}
//func main() {
//	price, num := 90, 6 // 定义 price 和 no,默认类型为 int
//	totalPrice := calculateBill(price, num)
//	fmt.Println("Total price is", totalPrice) // 打印到控制台上
//}


//package main
//
//import "fmt"
//
//func test(age int, num int) int {
//	totalnum := age * num
//	return totalnum
//}
//
//func main() {
//	var (
//		a     = 7850
//		n     = 12
//		total = test(a, n)
//	)
//	fmt.Println("total nume:", total)
//}


//多返回值，计算：    面积 = 长 * 宽    周长 = 2 * ( 长 + 宽 )
/*
package main

import (
    "fmt"
)

func rectProps(length, width float64)(float64, float64) {
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}

func main() {
    area, perimeter := rectProps(10.8, 5.6)
    fmt.Printf("Area %f Perimeter %f", area, perimeter)
}
*/

//_ 在 Go 中被用作空白符，可以用作表示任何类型的任何值,只调用函数的一个返回值
package main

import "fmt"

func rectProps(length, width float64) (float64, float64) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}

func main() {
	area, _ := rectProps(5.5, 2.5)
	fmt.Printf("Area %f\n", area)
}
