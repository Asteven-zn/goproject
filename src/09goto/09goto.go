package main

import "fmt"

func main() {
	//goto语句通过标签进行代码间的无条件跳转
	//goto语句可以在快速跳出循环、避免重复退出上有一定的帮助
	//Go语言中使用goto语句能简化一些代码的实现过程。 例如双层嵌套的for循环要退出时

	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			// 设置退出标签
	// 			goto breakTag
	// 		}
	// 		fmt.Printf("%v-%v\n", i, j)
	// 	}
	// }
	// return
	// 标签
	// breakTag:
	// 	fmt.Println("结束for循环")

	//不实用goto
	// var breakFlag bool
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			// 设置退出标签
	// 			breakFlag = true
	// 			break
	// 		}
	// 		fmt.Printf("%v-%v\n", i, j)
	// 	}
	// 	// 外层for循环判断
	// 	if breakFlag {
	// 		break
	// 	}
	// }

	var flag = false
	for i := 0; i <= 5; i++ {
		for j := 'A'; j <= 'C'; j++ {
			if j == 'C' {
				flag = true
				break
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			continue
		}
	}
}