package main
import "fmt"
func main(){
	// %v:值的默认格式表示
	// %d:表示为十进制

	//假如还有97天放假，问：XX个星期零XX天
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("还有%d个星期零%d天放假\n",week,day)

	//定义一个变量保存化氏温度,华氏温度转化摄氏度的公式：5/9*（华氏温度-100）
	//请求出华氏温度对应的摄氏温度
	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9.0 * (huashi - 100.0)
	fmt.Printf("%v化氏温度对应的摄氏温度=%v",huashi,sheshi)
}