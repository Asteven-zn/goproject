package main

import "fmt"

//流程控制调出for循环

func main() {
	//当i=5时就跳出for循环
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break //跳出for循环
		}
		fmt.Println(i)
	}
	fmt.Println("game over")

	//当i=5时就跳过for循环（不执行for循环内部的打印语句），继续下一次循环
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue //跳过条件，继续下一次循环
		}
		a := i + 10
		fmt.Println(a)
	}
	fmt.Println("game over")

	//跳出多层循环,break
	fmt.Println("############ 跳出多层循环 ############")
	for i := 0; i <= 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				break //跳出内层循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}

	//跳出多层循环,continue
	fmt.Println("############ 跳出多层循环 ############")
	for i := 0; i <= 10; i++ {
		for j := 'A'; j <= 'D'; j++ {
			if j == 'C' {
				continue //跳出内层循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
}
