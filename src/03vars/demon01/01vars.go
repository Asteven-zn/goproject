package main

import "fmt"

func main() {
	//格式1：定义一个变量，并且赋予变量类型
	var i int // var i int = 10
	i = 10
	fmt.Println("i = ", i)

	//格式2：定义一个变量，变量类型自动推导
	var num = 0.1
	fmt.Println("num = ", num)

	//格式3：定义一个变量，不写var
	numm := 2.15
	fmt.Println("numm = ", numm)

	//格式4：定义多个变量
	var id, name, age = 1, "tom", 18
	fmt.Println("id=", id, "name=", name, "age=", age)

	idd, namee, agee := 2, "jack", 29
	fmt.Println("id=", idd, "name=", namee, "age=", agee)

	//格式5：定义多个变量
	var (
		n1 = 1
		n2 = 2.15
		n3 = "steven"
	)
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
}
