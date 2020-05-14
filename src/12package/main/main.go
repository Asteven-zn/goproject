package main

import (
	"12package/utils"
	"fmt"
)

func main() {
	var n1 float64 = 1000
	var n2 float64 = 999
	var operator byte = '+'
	result := utils.Cal(n1, n2, operator)
	fmt.Println("result=", result)
}
