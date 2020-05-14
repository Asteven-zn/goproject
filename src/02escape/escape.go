package main
import "fmt"
func main()  {
	//\t：制表符
	//一个\：为转意思
	//\n：换行
	//\r：一个回车，后面的字符放到首行并覆盖相同的字符
	fmt.Println("tom\tjack")

	fmt.Println("hello\tworld")
	fmt.Println("C:\\usr\\local\\bin")
	fmt.Println("tom sad \"im am tom!\"")
	fmt.Println("天龙八部雪山飞狐\r张飞厉害")

	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t20\t北京\t昌平")

}