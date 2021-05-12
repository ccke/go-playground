package main

import (
	"fmt"
	"github.com/ccke/go-playground/study_01/homework_03/Service"
)

func main()  {
	// 1.初始化数据
	var companyName, employeeName, code, action, msg string
	var num, optionId, orderId int
	var res bool
	fmt.Print("开始初始化数据，请输入公司名、员工名、授予编号、授予数量：")
	fmt.Scan(&companyName, &employeeName, &code, &num)
	service := &Service.Service{}
	companyId, employeeId := service.InitData(companyName, employeeName, code, num)

	// 2.操作
	service.Login(companyId, employeeId)
	fmt.Println("已成功登录ESOP系统！")
	flag := true
	for flag {
		fmt.Println("-----请输入要执行的操作：0-订单展示，1-下单，2-改单，3-撤单，1000-订单后台成交，Q-退出-----")
		fmt.Scan(&action)
		switch action {
		case "0":
			service.ShowOrders()
		case "1":
			service.ShowOptions()
			fmt.Print("开始行权下单，请输入授予ID和行权数量：")
			fmt.Scan(&optionId, &num)
			orderId, res, msg = service.AddOrder(optionId, num)
			if res {
				fmt.Printf("行权下单操作成功，订单ID=%d，订单数量=%d\n", orderId, num)
			} else {
				fmt.Println(msg)
			}
		case "2":
			fmt.Print("开始行权改单，请输入订单ID和改单数量：")
			fmt.Scan(&orderId, &num)
			res, msg = service.ModifyOrder(orderId, num)
			if res {
				fmt.Printf("行权改单操作成功，订单ID=%d，订单数量=%d\n", orderId, num)
			} else {
				fmt.Println(msg)
			}
		case "3":
			fmt.Print("开始行权撤单，请输入订单ID：")
			fmt.Scan(&orderId)
			res, msg = service.CancelOrder(orderId)
			if res {
				fmt.Println("行权撤单操作成功。")
			} else {
				fmt.Println(msg)
			}
		case "1000":
			fmt.Print("开始后台成交，请输入订单ID：")
			fmt.Scan(&orderId)
			res, msg = service.FinishOrder(orderId)
			if res {
				fmt.Println("行权订单后台操作成功。")
			} else {
				fmt.Println(msg)
			}
		case "Q":
			fmt.Println("已退出ESOP系统，886~")
			flag = false
		default:
			fmt.Println("操作错误！")
		}
	}
}