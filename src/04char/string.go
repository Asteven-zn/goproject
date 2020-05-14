package main
import "fmt"
func main(){
	var (
		address string = "北京 110 西城 helloworld "
	)
	//输出字符串
	fmt.Println(address)
	
	//输出数值类型的默认值
	// %d：表示为十进制
	// %v：值的默认格式表示
	var (
		a int
		aa float32
		b float64
		c bool
		d string
	)
	fmt.Printf("%d\t%v\t%v\t%v\t%v",a,aa,b,c,d)

	//基本数据类型的相互转换
	var (
		i int32 = 100
		n1 float32 = float32(i) + 1.2
		n2 int8 = int8(i) + 2
		n3 int64 = int64(i) + 3
	)
		fmt.Printf("\ni=%v n1=%v n2=%v n3=%v",i,n1,n2,n3)
}