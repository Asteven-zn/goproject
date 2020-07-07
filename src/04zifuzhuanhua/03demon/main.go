package main

import "fmt"

//c%:转化位字符
func main() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()

	//字符串修改
	a1 := "红萝卜"
	a2 := []rune(a1) //把字符串强制转化为一个rune切片
	a2[0] = '白'
	fmt.Println(string(a2))

	//类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("f type %T\n", f)
}
