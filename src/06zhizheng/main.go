package main
import "fmt"
func main() {
	//基本数据类型在内存布局
	var i int = 10
	// i 的地址是什么，&i
	fmt.Println("i的地址=", &i)

	//下面的var ptr *int = &i
	//1. ptr 是一个指针变量
	//2. ptr 的类型 *int
	//3. ptr 本身的字&i
	var ptr *int = &i
	fmt.Printf("ptr 存储的是i的地址=%v\n",ptr)
	fmt.Printf("ptr 的类型=%T\n",ptr)
	fmt.Printf("ptr 的地址=%v\n",&ptr)
	fmt.Printf("prt 指向的值=%v",*ptr)
}