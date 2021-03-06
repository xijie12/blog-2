package main

import (
	"fmt"
	"time"
)

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}

func bufferedChannel() {
	c := make(chan int, 3)
	//read(c)
	c <- 1
	c <- 2
	c <- 3
	for i := 0; i < 10; i++ {
		c <- i
	}
	time.Sleep(time.Second)
}

func channelClose() {
	c := make(chan int, 3)
	read(c)
	c <- 1
	c <- 2
	c <- 3
	c <- 5
	//close(c)

	time.Sleep(time.Second)
}

func read(ch <-chan int) {

	go func() {
		for  n := range ch {
			//n, ok := <-ch
			//if !ok {
			//	break
			//}
			fmt.Println(n)
		}
	}()
}

// 外部只能发数据
func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker recevice id=%d chan=%c\n", id, <-c)
		}
	}()

	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Second)

}
