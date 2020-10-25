package utils

import (
	"fmt"
)

type FamilyAccount struct {
	//声明必要的字段
	//声明一个变量，保存接收用户输入的选项
	key string
	//声明一个变量，控制是否退出for循环
	loop bool
	//定义账户的余额
	balance float64
	//每次收入的金额
	money float64
	//每次收支的说明
	note string
	//定义个字段，记录是否有收支的行为
	flag bool
	//收支的详情，使用字符串来记录
	//当有收支时，只需要对details进行拼接处理即可
	details string
}

//编写一个工厂模式的构造方法，返回一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key : "",
		loop : true,
		balance : 10000.0,
		money : 0.0,
	    note : "",
		flag : false,
		details : "收支\t账户金额\t收支金额\t说	明",
	}
}


//将显示明细写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) showDetails() {
	fmt.Println("*****************当前收支明细*****************")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细明细，，，来一笔吧！")
	}
}

//将登记收入写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) incom() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money //修改账户余额
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	//将本次收入情况，拼接到details变量
	//收入   11000      1000    有人发红包
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//将登记登记支出写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) pay() {
	fmt.Println("*****************登记支出********************")
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	//这里做一个支出判断，当支出大于余额退出
	if this.money > this.balance {
		fmt.Println("账户余额不足")
		//break
	}
	this.balance -= this.money //修改账户余额
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	//将本次支出情况，拼接到details变量
	//支出   11000      1000    吃饭
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//将退出系统写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你输入的有误，请重新输入 y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}

//给改结构体绑定相应的方法
//显示主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n-------------------家庭收支记账软件-------------------")
		fmt.Println("                   1 收支明细")
		fmt.Println("                   2 登记收入")
		fmt.Println("                   3 登记支出")
		fmt.Println("                   4 退出软件")
		fmt.Println("请选择功能(1-4)")
		fmt.Scanln(&this.key)
		switch this.key {
			case "1":
				this.showDetails()
			case "2":
				this.incom()
			case "3":
				this.pay()
			case "4":
				this.exit()
			default:
				fmt.Println("*****************请收入正确选项***************")
		}
		//loop初始值为true，选4的时候为false
		//if判断调价是用！号取loop的反就是false，这个时候满足条件继续往下执行
		if !this.loop {
			break
		}
	}	
} 