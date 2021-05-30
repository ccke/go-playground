package test

func Sort(data []int) {
	size := len(data)
	for i := 0; i < size - 1; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] > data[j] {
				temp := data[i]
				data[i] = data[j]
				data[j] = temp
			}
		}
	}
}
