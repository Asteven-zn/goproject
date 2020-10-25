package main

import "fmt"

func main()  {

	//声明一个变量，保存接收用户输入的选项
	key := ""
	//声明一个变量，控制是否退出for循环
	loop := true

	//定义账户的余额
	balance := 10000.0
	//每次收入的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//收支的详情，使用字符串来记录
	//当有收支时，只需要对details进行拼接处理即可
	details := "收支\t账户金额\t收支金额\t说	明"

	//显示主菜单
	for {
		fmt.Println("\n-------------------家庭收支记账软件-------------------")
		fmt.Println("                   1 收支明细")
		fmt.Println("                   2 登记收入")
		fmt.Println("                   3 登记支出")
		fmt.Println("                   4 退出软件")
		fmt.Println("请选择功能(1-4)")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("*****************当前收支明细*****************")
			fmt.Println(details)
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money  //修改账户余额
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			//将本次收入情况，拼接到details变量
			//收入   11000      1000    有人发红包
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
		case "3":
			fmt.Println("*****************登记支出********************")
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			//这里做一个支出判断，当支出大于余额退出
			if money > balance {
				fmt.Println("账户余额不足")
				break
			}
			balance -= money  //修改账户余额
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			//将本次支出情况，拼接到details变量
			//支出   11000      1000    吃饭
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
		case "4":
			fmt.Println("你确定要退出吗？y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你输入的有误，请重新输入 y/n")
			}
			loop = false
		default:
			fmt.Println("*****************请收入正确选项***************")
		}
		//loop初始值为true，选4的时候为false
		//if判断调价是用！号取loop的反就是false，这个时候满足条件继续往下执行
		if !loop {
			break
		}
	}
	fmt.Println("你退出了家庭记账软件.........")
}