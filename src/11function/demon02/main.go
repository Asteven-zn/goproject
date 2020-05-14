package main

import "fmt"

//getsum 计算两个数的和
func getsum(n1 int , n2 int) int {
	sum := n1 + n2
	fmt.Println("getsum sum = ", sum)
	return sum
}

//计算两个数的和与差
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}


//getsum 计算两个数的和调用
func main() {
	sum := getsum(10, 20)
	fmt.Println("main sum = ", sum)

//计算两个数的和与差调用
	res1, res2 := getSumAndSub(1, 2)
	fmt.Printf("res1=%v res2=%v\n", res1, res2)

//忽略某个值，使用 _ 符号
	_, res3 := getSumAndSub(1, 2)
	fmt.Printf("res3=%v\n", res3)
}