package main

import "fmt"

//基本格式
func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//变种1
	var i = 5
	for ; i <= 10; i++ {
		fmt.Println(i)
	}

	//变种2
	var a = 5
	for a <= 10 {
		fmt.Println(a)
		a++
	}

	//range遍历
	s := "hello深圳"
	for id, date := range s {
		fmt.Printf("[%d]:%c\n", id, date)
	}

	//switch_case基本格式
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}

	//变种1，case可以判断多种情况
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	//变种2，case可以时表达式
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	//fallthrough语法可以执行满足条件的case的下一个case
	m := "hello go"
	switch {
	case m == "hello go":
		fmt.Println("hello go")
		fallthrough
	case m == "love you":
		fmt.Println("love you")
	case m == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
