package main
import (
	"fmt"
	_ "unsafe"
	"strconv"
)

//演示golang中基本数据转换成string使用
func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string //空的str
	
	//使用第一种方式来转换 fmt.Sprintf 的方法
	//%q:该值对应的字符串值，用双引号括起来
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str,str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str,str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str,str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str=%q\n\n", str,str)

	//第二种方式 strconv 函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10) //10代表转成10进制
	fmt.Printf("str type %T str=%q\n", str,str)

	// strconv.FormatFloat(num4,'f', 10, 64)
	// 说明：'f'格式   10：表示小数位保留10位    64：表示这个小数是float64位
	str = strconv.FormatFloat(num4,'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str,str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str,str)
}