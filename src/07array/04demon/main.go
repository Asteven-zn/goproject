package main

import "fmt"

//多维数组，元素数量只有外层数组可以用“...”自动推导

func array(a [3] [2] int) {
	a[0][0] = 100
}

func main()  {
	x := [3][2]int{
		{1, 11},
		{2, 22},
		{3, 33},
	}
	array(x)
	fmt.Println(x)

	y := x
	y [1][1] = 100
	fmt.Println(x)
}