package main

import "fmt"

func test() (int, string) {
	return 29,"cloud"
}

func main()  {
//匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
	age, _ := test()
	_, name := test()
	age1, name1 := test()
	fmt.Println("age=", age)
	fmt.Println("name=", name)
	fmt.Println("age1=", age1 ,"name1=", name1)
}
