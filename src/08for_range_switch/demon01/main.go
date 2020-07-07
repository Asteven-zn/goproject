package main

import "fmt"

// func gotoDemo1() {
// 	var breakFlag bool
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			if j == 2 {
// 				// 设置退出标签
// 				breakFlag = true
// 				break
// 			}
// 			fmt.Printf("%v-%v\n", i, j)
// 		}
// 		// 外层for循环判断
// 		if breakFlag {
// 			break
// 		}
// 	}
// }

// a := func gotoDemo1()

func main() {
	var a = [...]string{"北京", "上海", "深圳"}
	for index, value := range a {
		fmt.Println(index, value)
	}
}
