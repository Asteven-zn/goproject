package main

import "fmt"

func main() {
	//ping 192.168.1-5(除了.3).1-3的ip
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue //跳过条件，继续下一次循环
		}
		for j := 1; j < 5; j++ {
			if j == 4 {
				break
			}
			fmt.Printf("ping 192.168.%v.%v\n", i, j)
		}
	}
	fmt.Println("###### test #######")
	//ping 192.168.1-5.1-3的ip
	for a := 1; a <= 5; a++ {
		for b := 1; b < 5; b++ {
			if b == 4 {
				break
			}
			fmt.Printf("ping 192.168.%v.%v\n", a, b)
		}
	}
}
