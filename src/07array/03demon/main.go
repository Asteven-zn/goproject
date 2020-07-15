package main

import "fmt"

//数组是指类型，调用时是复制了一份

func array1(a [3]int) {
	a[0] = 100
}

func main() {
	x := [3]int{1, 2, 3}
	fmt.Println(x)
    array1(x)
	fmt.Println(x)
}