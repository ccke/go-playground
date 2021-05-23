package main

import (
	"fmt"
	"time"
)

func main() {
	chList := []chan int{make(chan int), make(chan int), make(chan int), make(chan int)}

	for i := 0; i < 4; i++ {
		go func(ch chan int, nextCh *chan int, no int) {
			for {
				token := <-ch
				fmt.Println("I'm goroutine ", no)
				time.Sleep(time.Second)
				*nextCh <- token
			}
		}(chList[i], &chList[(i+1)%4], i+1)
	}
	chList[0] <- 1

	select {}
}