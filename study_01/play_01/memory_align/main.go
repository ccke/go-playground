package main

import (
	"fmt"
	"unsafe"
)

type p1 struct {
	a byte
	b int32
	c int8
	d int64
	e bool
}

type p2 struct {
	a byte
	c int8
	e bool
	b int32
	d int64
}

// 内存对齐测试
func main()  {
	// 不紧凑
	p1 := p1{
		'a',
		32,
		8,
		64,
		true,
	}
	fmt.Printf("调整前 size: %d, align: %d\n", unsafe.Sizeof(p1), unsafe.Alignof(p1))

	// 紧凑
	p2 := p2{
		'a',
		8,
		true,
		32,
		64,
	}
	fmt.Printf("调整后 size: %d, align: %d\n", unsafe.Sizeof(p2), unsafe.Alignof(p2))

	// 指针大小
	fmt.Printf("指针 size: %d\n", unsafe.Sizeof(&p2))
}
