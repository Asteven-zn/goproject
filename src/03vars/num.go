package main
import (
	"fmt"
	"unsafe"
)
func main() {
	//%T 输出值的类型，表示为十进制
	//%d 输出数据所占用的字节数
	var n1 = -1000000
	fmt.Printf("%T \n%d",n1,unsafe.Sizeof(n1))
	n2 := 20
	fmt.Println("\nn2 =",n2)

	// var c1 byte = 'a'
	// var c2 byte = '3'
	// %c值对应的unicode码值
	var (
		c1 byte = 'a'
		c2 byte = '3'
	)
	fmt.Printf("%c\n%c", c1, c2)
}