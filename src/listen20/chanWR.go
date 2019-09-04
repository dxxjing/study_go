package main

import (
	"fmt"
	"sync"
	"time"
)

//该协程只负责写入 不能读取 在写入的函数中关闭channel
func writeChan(ch chan<- int,wg *sync.WaitGroup){
	for i := 0; i < 10; i++{
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)

	wg.Done()
}

//该协程负责读取 不能写入
func readChan(ch <-chan int,wg *sync.WaitGroup){
	for v := range ch {
		fmt.Println(v)
	}

	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int,10)

	go writeChan(ch,&wg)
	go readChan(ch,&wg)

	wg.Wait()
	fmt.Println("all goroutine done")
}
