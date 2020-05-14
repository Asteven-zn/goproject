package main
import "fmt"
func main(){
	//演示逻辑运算符的使用
	var age int = 40
	//&&:与运算，两个条件都必须成立
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	//||:或运算，两个条件成立一个即可	
	if age > 30 || age <= 40 {
		fmt.Println("ok2")
	}
	
	//!:非如果条件为true，则为fales。如果条件为fales，则为true
	var age1 int = 20
	if !(age1 > 30) {
		fmt.Println("ok3")
	}

	//两个变量，a和b，要求将其进行交换，但是不允许使用中间变量，最终打印结果
	var a int = 10
	var b int = 30
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a=%v b=%v\n",a,b)

	//求两个数的最大值
	var n1 int = 1
	var n2 int = 2
    var max int
	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}
	fmt.Printf("max=%v\n",max)

	//求3个数的最大值
	var n3 int = 3
	if n3 > max {
		max = n3
	}
	fmt.Printf("三个数最大的数是%v",max)
}