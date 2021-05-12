package Service

import (
	"fmt"
	"github.com/ccke/go-playground/study_01/homework_03/models"
	"math/rand"
)

// Data 假设只有一个公司、一个员工、一条授予
type Data struct {
	Company models.Company
	Employee models.Employee
	Option models.Option
	OrderList []models.Order
}

type Service struct {
	CompanyId int
	EmployeeId int
	Data Data
}

func (service *Service) InitData(companyName string, employeeName string, code string, num int) (int, int) {
	// 1.初始化默认信息
	if companyName == "" {
		companyName = "默认公司"
	}
	if employeeName == "" {
		employeeName = "默认员工"
	}
	if code == "" {
		code = "默认授予编号"
	}
	if num == 0 {
		num = 10000
	}

	// 2.创建公司信息
	companyId := rand.Int()
	service.Data.Company = models.Company{
		Id:   companyId,
		Name: companyName,
	}

	// 3.创建员工信息
	employeeId := rand.Int()
	service.Data.Employee = models.Employee{
		Id:        employeeId,
		CompanyId: companyId,
		Name:      employeeName,
	}

	// 4.创建授予信息
	optionId := rand.Int()
	service.Data.Option = models.Option{
		Id:            optionId,
		CompanyId:     companyId,
		EmployeeId:    employeeId,
		Code:          code,
		Num:           num,
		UnExerciseNum: num,
	}

	return companyId, employeeId
}

func (service *Service) Login(companyId int, employeeId int)  {
	service.CompanyId = companyId
	service.EmployeeId = employeeId
}

// 下单
func (service *Service) AddOrder(optionId int, num int) (int, bool, string) {
	// 1.检验是否有权限
	if service.CompanyId != service.Data.Option.CompanyId || service.EmployeeId != service.Data.Option.EmployeeId {
		return 0, false, "无操作权限"
	}

	// 2.检验参数是否正确
	if optionId <= 0 || service.Data.Option.Id != optionId || num <= 0 {
		return 0, false, "参数不正确"
	}

	// 3.检验是否够下单
	if service.Data.Option.Num < num {
		return 0, false, "行权数量超出"
	}
	exercisedNum := 0
	for _, order := range service.Data.OrderList {
		if order.Status != models.Canceled {
			exercisedNum += order.Num
		}
	}
	if service.Data.Option.Num < num + exercisedNum {
		return 0, false, "行权数量超出"
	}

    // 4.发起行权
    service.Data.Option.UnExerciseNum -= num
    orderId := rand.Int()
    service.Data.OrderList = append(service.Data.OrderList, models.Order{
		Id:         orderId,
		CompanyId:  service.CompanyId,
		EmployeeId: service.EmployeeId,
		OptionId:   optionId,
		Num:        num,
		Status:     models.Init,
	})
    return orderId, true, ""
}

// 后台订单成交
func (service *Service) FinishOrder(orderId int) (bool, string) {
	// 1.检验是否有权限
	if service.CompanyId != service.Data.Option.CompanyId || service.EmployeeId != service.Data.Option.EmployeeId {
		return false, "无操作权限"
	}

	// 2.检验参数是否正确
	if orderId <= 0 {
		return false, "参数不正确"
	}

	// 3.后台订单成交
	for k, order := range service.Data.OrderList {
		if order.Id == orderId {
			if order.Status != models.Init {
				return false, "订单已完成，不可后台操作"
			}
			service.Data.OrderList[k].Status = models.Finished
			return true, ""
		}
	}
	return false, "订单信息不存在"
}

// 改单
func (service *Service) ModifyOrder(orderId int, num int) (bool, string) {
	// 1.检验是否有权限
	if service.CompanyId != service.Data.Option.CompanyId || service.EmployeeId != service.Data.Option.EmployeeId {
		return false, "无操作权限"
	}

	// 2.检验参数是否正确
	if orderId <= 0 || num <= 0 {
		return false, "参数不正确"
	}

	// 3.发起改单
	for k, order := range service.Data.OrderList {
		if order.Id == orderId {
			if order.Status != models.Init {
				return false, "订单已完成，不允许改单"
			}
			if num - order.Num > service.Data.Option.UnExerciseNum {
				return false, "改单数量超出"
			}
			service.Data.OrderList[k].Num = num
			service.Data.Option.UnExerciseNum -= num - order.Num
			return true, ""
		}
	}
	return false, "订单信息不存在"
}

// 撤单
func (service *Service) CancelOrder(orderId int) (bool, string) {
	// 1.检验是否有权限
	if service.CompanyId != service.Data.Option.CompanyId || service.EmployeeId != service.Data.Option.EmployeeId {
		return false, "无操作权限"
	}

	// 2.检验参数是否正确
	if orderId <= 0 {
		return false, "参数不正确"
	}

	// 3.发起撤单
	for k, order := range service.Data.OrderList {
		if order.Id == orderId {
			if order.Status != models.Init {
				return false, "订单已完成，不允许撤单"
			}
			service.Data.OrderList[k].Status = models.Canceled
			service.Data.Option.UnExerciseNum += order.Num
			return true, ""
		}
	}
	return false, "订单信息不存在"
}

// 订单记录
func (service Service) ShowOptions() {
	option := service.Data.Option
	fmt.Printf("Option授予记录：授予ID=%d，授予数量=%d，可行权数量=%d\n", option.Id, option.Num, option.UnExerciseNum)
}

// 订单记录
func (service Service) ShowOrders() {
	fmt.Printf("订单记录如下：\n")
	for k, order := range service.Data.OrderList {
		fmt.Printf("%d、订单ID：%d，行权数量：%d，订单状态：%s\n", k + 1, order.Id, order.Num, order.Status)
	}
}
