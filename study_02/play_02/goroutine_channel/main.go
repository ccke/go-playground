package main

import (
	"fmt"
	"sync"
)

func send(ch chan int)  {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("开始放入：", i)
	}
	close(ch)
}

func receive(ch1, ch2 chan int)  {
	defer wg.Done()
	for {
		i, ok := <- ch1
		if ok {
			fmt.Println("开始取出：", i)
			ch2 <- i * i
		} else {
			break
		}
	}
}

func compute(ch chan int)  {
	close(ch)
	for v := range ch {
		fmt.Println("计算结果：", v)
	}
}

var wg sync.WaitGroup
var once sync.Once

func main()  {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	wg.Add(3)
	// 生产者
	go send(ch1)
	// 消费者
	go receive(ch1, ch2)
	go receive(ch1, ch2)
	wg.Wait()
	// 计算
	compute(ch2)
}
