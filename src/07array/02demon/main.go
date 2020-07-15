package main

import (
	"fmt"
)

func main()  {
	var a = [...]string{"北京", "上海", "深圳"}

	fmt.Println("###方法1：for循环遍历")
	for i :=0; i < len(a); i++{
		fmt.Println(i, a[i])
	}

	fmt.Println("###方法2：for range遍历")
	for index, value := range a {
		fmt.Println(index, value)
	}
}
