package main

import (
	"fmt"
	"math/rand"
)



func main() {
	s := make([]int, 0, 16)
	for i := 0; i < 16; i++ {
		s = append(s, rand.Intn(100))
	}
	fmt.Println("排序前：", s)
	quickSort(s)
	fmt.Println("排序后：", s)
}
