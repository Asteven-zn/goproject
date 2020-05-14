package utils

//init函数，通常可以在init函数中完成初始化工作，在main函数执行前就被调用
import "fmt"
var Age int
var Name string

//Age 和 Name 是全局变量，我们需要在main.go 使用
//但是我们需要初始化Age 和 Name

func init() {
	fmt.Println("utils 包的 inti()....")
	Age = 100
	Name = "tom~"
}