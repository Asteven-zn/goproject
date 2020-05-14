package main

import (
	"fmt"
)

func sum(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1=", n1) //def 3. ok1 n1 = 10
	defer fmt.Println("ok2 n2=", n2) //def 2. ok1 n2 = 20

	res := n1 + n2  // res = 30
	fmt.Println("ok3 res=", res)  //1. ok3 res = 30

	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)  // 4. res = 30
}