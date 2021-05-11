package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1.公司7个员工在这一年中每人的收益
	incomeMap := []float32{5.5, 10.7, 6.22, 20.1, 8.1, 12.2, 30.2}
	fmt.Printf("未排序前数据：%v \n", incomeMap)

	// 2.根据员工的收益进行由高到低排序
	nameArray := rank(incomeMap)
	fmt.Printf("收入排序后的数据：%v \n", nameArray)

	// 3. 统计所有人累计收益
	totalIncome := total(incomeMap)
	fmt.Printf("所有员工累计收益：%v \n", totalIncome)
}

// rank 对7名员工的收益进行由高到低排序,并返回排序后数组
func rank(data []float32) []float32 {
	size := len(data)
	for i := 0; i < size - 1; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] < data[j] {
				temp := data[i]
				data[i] = data[j]
				data[j] = temp
			}
		}
	}
	return data
}

// total 统计员工累计收益
func total(data []float32) float32 {
	var sum float32 = 0.0
	for _, value := range data {
		sum += value
	}
	result, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", sum),64)
	return (float32)(result)
}
