package main

// 冒泡排序
func bubblingSort(data []int) {
	size := len(data)
	if size <= 1 {
		return
	}
	swap := false
	for i := size-1; i > 0; i--  {
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
	for i := 0; i < size - 1; i++  {
		for j := i+1; j < size; j++ {
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
	for i := 1; i < size; i++  {
		base := data[i]
		j := i - 1
		for ; j >= 0 && data[j] > base; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = base
	}
}

// 归并排序
// https://zhuanlan.zhihu.com/p/42586566
// https://blog.csdn.net/alzzw/article/details/98100378
func mergeSort(data []int)  {
	if len(data) <= 1 {
		return
	}
	mid := (len(data) -1)/2
	mergeSort(data[:mid])
	mergeSort(data[mid+1:])
	mergeArray(data, mid)
}


func mergeArray(data []int, mid int) {
	size := len(data)
	var temp = make([]int, size)
	k, i, j := 0, 0, mid+1
	for i <= mid && j < size {
		if data[i] < data[j] {
			temp[k] = data[i]
			i++
		} else {
			temp[k] = data[j]
			j++
		}
		k++
	}
	for i <= mid {
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
	// 分区
	base := data[0]
	l, r := 0, len(data)-1
	for i := 1; i <= r; {
		if data[i] > base {
			data[i], data[r] = data[r], data[i]
			r--
		} else {
			data[i], data[l] = data[l], data[i]
			l++
			i++
		}
	}
	// 递归排序
	quickSort(data[:l])
	quickSort(data[l+1:])
}
