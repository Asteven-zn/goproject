package main
import (
	"familyaccountapp2/utils"
	"fmt"
)

func main() {
	fmt.Println("这个是面向对象方法的完成！")
	utils.NewFamilyAccount().MainMenu()
}