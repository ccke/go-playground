package main

import (
	"math"
	"sort"
)

// 冒泡排序
func bubblingSort(data []int) {
	size := len(data)
	if size <= 1 {
		return
	}
	swap := false
	for i := size - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
}

// 选择排序
func selectionSort(data []int) {
	size := len(data)
	if size <= 1 {
		return
	}
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] > data[j] {
				data[j], data[i] = data[i], data[j]
			}
		}
	}
}

// 插入排序
func insertSort(data []int) {
	size := len(data)
	if size <= 1 {
		return
	}
	for i := 1; i < size; i++ {
		base := data[i]
		j := i - 1
		for ; j >= 0 && data[j] > base; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = base
	}
}

// 归并排序
func mergeSort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := len(data) / 2
	mergeSort(data[:mid])
	mergeSort(data[mid:])
	mergeArray(data, mid)
}

// 合并数组
func mergeArray(data []int, mid int) {
	size := len(data)
	var temp = make([]int, size)
	k, i, j := 0, 0, mid
	for i < mid && j < size {
		if data[i] < data[j] {
			temp[k] = data[i]
			i++
		} else {
			temp[k] = data[j]
			j++
		}
		k++
	}
	for i < mid {
		temp[k] = data[i]
		i++
		k++
	}
	for j < size {
		temp[k] = data[j]
		j++
		k++
	}
	for i := 0; i < size; i++ {
		data[i] = temp[i]
	}
}

// 快速排序
func quickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	l := partition(data)
	quickSort(data[:l])
	quickSort(data[l+1:])
}

// 分区
func partition(data []int) int {
	base := data[0]
	l, r := 0, len(data)-1
	for l < r {
		for l < r && data[r] >= base {
			r--
		}
		data[l] = data[r]
		for l < r && data[l] <= base {
			l++
		}
		data[r] = data[l]
	}
	data[l] = base
	return l
}

// 堆排序
func heapSort(data []int) {
	size := len(data)
	if size <= 1 {
		return
	}
	// 构建堆顶
	for i := size/2 - 1; i >= 0; i-- {
		adjustHeap(data, i, size)
	}
	// 调整堆结构+交换堆顶元素与末尾元素
	for i := size - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		adjustHeap(data, 0, i)
	}
}

// 堆调整
func adjustHeap(data []int, start int, end int) {
	temp := data[start]
	for i := start*2 + 1; i < end; i = i*2 + 1 {
		if i+1 < end && data[i] < data[i+1] {
			i++
		}
		if data[i] > temp {
			data[start] = data[i]
			start = i
		} else {
			break
		}
	}
	data[start] = temp
}

// 希尔排序
// 常见的增量序列：
// 1.最初Donald Shell提出的增量，即折半降低直到1。据研究，使用希尔增量，其时间复杂度还是O(n2)。
// 2.Hibbard增量：{1, 3, ..., 2k-1}，该增量序列的时间复杂度大约是O(n1.5)。
// 3.Sedgewick增量：(1, 5, 19, 41, 109,...)，其生成序列或者是94i - 92i + 1或者是4i - 3*2i + 1。
func shellSort(data []int) {
	size := len(data)
	if size <= 1 {
		return
	}
	for delta := size / 2; delta >= 1; delta /= 2 {
		for i := delta; i < size; i++ {
			for j := i; j >= delta && data[j] < data[j-delta]; j -= delta {
				data[j], data[j-delta] = data[j-delta], data[j]
			}
		}
	}
}

// 计数排序
func countSort(data []int) {
	size := len(data)
	if size <= 1 {
		return
	}
	// 找出数组最大最小值
	min, max := math.MaxInt8, math.MinInt8
	for _, v := range data {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	// 计数数组
	c := make([]int, max-min+1)
	for _, v := range data {
		c[v-min]++
	}
	// 将值替换到原数组
	index := 0
	for i, v := range c {
		for v > 0 {
			data[index] = i + min
			v--
			index++
		}
	}
}

// 桶排序
func bucketSort(data []int) {
	size := len(data)
	if size <= 1 {
		return
	}
	// 找出数组最大最小值
	min, max := math.MaxInt8, math.MinInt8
	for _, v := range data {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	// 桶
	bucketNum := (max-min)/size+1
	bucket := make([][]int, bucketNum)
	for _, v := range data {
		index := (v-min)/size
		bucket[index] = append(bucket[index], v)
	}
	// 对每个桶进行排序，将值替换到原数组
	index := 0
	for _, items := range bucket {
		// 底层用的快排，因此不稳定
		sort.Ints(items)
		for _, v := range items {
			data[index] = v
			index++
		}
	}
}

// 基数排序
func radixSort(data []int)  {
	size := len(data)
	if size <= 1 {
		return
	}
	// 找出数组最大值
	max:= math.MinInt8
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	// 按位数排序
	for exp := 1; max/exp > 0; exp *= 10 {
		bucket := make([][]int, 10)
		for i := 0; i < size; i++ {
			index := (data[i]/exp)%10
			bucket[index] = append(bucket[index], data[i])
		}
		index := 0
		for _, items := range bucket {
			for _, v := range items {
				data[index] = v
				index++
			}
		}
	}
}