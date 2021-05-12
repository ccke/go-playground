package models

// Order 订单信息
type Order struct {
	Id int
	CompanyId int
	EmployeeId int
	OptionId int
	Num int          // 下单数量
	Status string    // 订单状态：Init、Canceled、Finished
}

const (
	Init string = "行权中"
	Canceled string = "已撤单"
	Finished string = "已成交"
)
