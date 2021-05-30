package test

import (
	"fmt"
	"reflect"
	"testing"
)

// 1.单元测试
func TestSort(t *testing.T)  {
	// 定义一个测试用例类型
	type test struct {
		input []int
		want  []int
	}
	// 定义一个存储测试用例的切片
	tests := map[string]test{
		"顺序错乱": {input: []int{1,4,3,5,6,2}, want: []int{1,2,3,4,5,6}},
		"顺序": {input: []int{1,2,3,4,5,6}, want: []int{1,2,3,4,5,6}},
		"倒序": {input: []int{6,5,4,3,2,1}, want: []int{1,2,3,4,5,6}},
	}
	// 遍历切片，逐一执行测试用例
	var got []int
	for name, tc := range tests {
		got = tc.input
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			Sort(got)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// 2.性能测试
func benchmarkSort(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Sort([]int{1,4,3,5,6,2})
	}
}

func BenchmarkSort1(b *testing.B)  {
	benchmarkSort(b)
}

func BenchmarkSort10(b *testing.B)  {
	benchmarkSort(b)
}

func BenchmarkSort100(b *testing.B)  {
	benchmarkSort(b)
}

// 3.示例函数
func ExampleSort() {
	data := []int{1,4,3,5,6,2}
	Sort(data)
	fmt.Println(data)
	// Output:
	// [1 2 3 4 5 6]
}