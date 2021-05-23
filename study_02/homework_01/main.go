package main

import (
	"fmt"
	"github.com/ccke/go-playground/study_02/homework_01/models"
	"github.com/ccke/go-playground/study_02/homework_01/services"
)

func main() {
	var list []models.Employee
	for i := 1; i < 2; i++ {
		list = append(list, models.Employee{
			Id:        i,
			CompanyId: i * 10,
			Name:      fmt.Sprintf("name%v", i),
			Uid:       fmt.Sprintf("%v", i*100),
			IdCode:    fmt.Sprintf("%v", i*1000),
		})
	}

	kms := services.NewKmsServer()
	if err := kms.Encrypt(list); err == nil {
		fmt.Println("编码后的数据为:", list)
		if err = kms.Decrypt(list); err == nil {
			fmt.Println("解码后的数据为:", list)
		} else {
			fmt.Println("解码失败")
		}
	} else {
		fmt.Println("编码失败")
	}
}
