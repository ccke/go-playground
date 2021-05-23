package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Age int    `json:"age"`
}

func main() {
	stu1 := student{
		Name:  "测试",
		Age: 24,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
}
