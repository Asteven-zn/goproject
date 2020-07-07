package main

import (
	"fmt"
)

func main()  {
	fmt.Println("###定义数组格式：var 数组变量名 [元素数量]T")
	var testarray [3]int
	var numarray = [3]int{1,2}
	var cityarray = [3]string{"北京", "上海", "深圳"}
	fmt.Println(testarray)
	fmt.Println(numarray)
	fmt.Println(cityarray)

	fmt.Println("###让编译器根据初始值的个数自行推断数组的长度")
	var array1 [3]int
	var array2 = [...]int{1,2}
	var array3 = [...]string{"北京", "上海","深圳"}

	fmt.Printf("array1 = %d, type of array1 :%T\n", array1, array1)
	fmt.Printf("array2 = %d, type of array2 :%T\n", array2, array2)
	fmt.Printf("array3 = %v, type of array3 :%T\n", array3, array3)
}
