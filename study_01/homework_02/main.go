package main

import (
	"fmt"
	"github.com/ccke/go-playground/study_01/homework_02/grant"
)

func main()  {
	// 1.创建DB
	var dbName string
	fmt.Print("开始创建数据库，请输入数据库名：")
	fmt.Scan(&dbName)
	db := &grant.PersonDb{
		Name: dbName,
	}

	// 2.插入数据
	var no string
	var count int
	fmt.Print("开始插入数据，请输入员工编号和激励数量：")
	fmt.Scan(&no, &count)
	res := db.Insert(grant.Person{
		No:    no,
		Count: count,
	})
	if res {
		fmt.Println("插入成功！")
		db.Show()
	} else {
		fmt.Println("插入失败！")
	}

	// 3.查找数据
	fmt.Print("开始查找数据，请输入员工编号：")
	fmt.Scan(&no)
	str := db.FindByNo(no)
	fmt.Println("查找结果：", str)

	// 4.更新数据
	fmt.Print("开始更新数据，请输入员工编号和激励数量：")
	fmt.Scan(&no, &count)
	res = db.Update(grant.Person{
		No:    no,
		Count: count,
	})
	if res {
		fmt.Println("更新成功！")
		db.Show()
	} else {
		fmt.Println("更新失败！")
	}

	// 5.删除数据
	fmt.Print("开始删除数据，请输入员工编号：")
	fmt.Scan(&no)
	res = db.DeleteByNo(no)
	if res {
		fmt.Println("删除成功！")
		db.Show()
	} else {
		fmt.Println("删除失败！")
	}
}