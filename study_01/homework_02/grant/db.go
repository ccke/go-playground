package grant

import (
	"encoding/json"
	"fmt"
)

// 用户
type Person struct {
	No    string // 员工编号
	Count int    // 激励数量
}

// Db接口
type Db interface {
	Insert(data interface{}) bool
	Update(data interface{}) bool
	FindByNo(no string) string
	DeleteByNo(no string) bool
}

// 用户Db
type PersonDb struct {
	Name string   // Db名称
	Count int     // 数据条数
	Data []Person // 数据
}

func (db *PersonDb) Insert(data Person) bool {
	// 1.规则校验：编号不为空、数量>0
	if data.No == "" || data.Count <= 0 {
		return false
	}

	// 2.逻辑校验：编号不重复
	for _, item := range db.Data {
		if data.No == item.No {
			return false
		}
	}

	// 3.数据插入逻辑
	db.Data = append(db.Data, data)
	db.Count++
	return true
}

func (db *PersonDb) Update(data Person) bool {
	// 1.规则校验：编号不为空、数量>0
	if data.No == "" || data.Count <= 0 {
		return false
	}

	// 2.数据更新逻辑
	for k, item := range db.Data {
		if data.No == item.No {
			db.Data[k] = data
			return true
		}
	}
	return false
}

func (db PersonDb) FindByNo(no string) string {
	// 数据查找逻辑
	for _, item := range db.Data {
		if no == item.No {
			str, _ := json.Marshal(item)
			return string(str)
		}
	}
	return ""
}

func (db *PersonDb) DeleteByNo(no string) bool {
	// 数据删除逻辑
	for k, item := range db.Data {
		if no == item.No {
			db.Data = append(db.Data[:k], db.Data[k+1:]...)
			db.Count--
			return true
		}
	}
	return false
}

func (db PersonDb) Show() {
	fmt.Printf("数据库 %s 共有 %d 条数据，具体如下：\n", db.Name, db.Count)
	for k, item := range db.Data {
		fmt.Printf("%d、员工编号：%s，激励数量：%d\n", k + 1, item.No, item.Count)
	}
}
