package main

import "fmt"

func main() {
	var age int = 40
	//逻辑与运算
	if age > 0 && age <= 20 {
		fmt.Println("未读大学")
	} else if age > 20 && age < 25 {
		fmt.Println("读大学")
	} else {
		fmt.Println("已经毕业了")
	}

	//逻辑或运算
	if age > 50 || age < 20 {
		fmt.Println("老人或小孩专座")
	} else if age > 20 || age < 50 {
		fmt.Println("普通座")
	} else {
		fmt.Println("无座")
	}
	//逻辑非运算
	if age > 30 {
		fmt.Println("应该成家了")
	}
	if !(age > 30) {
		fmt.Println("还不着急成家")
	}
}
