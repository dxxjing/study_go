package main

import (
	"fmt"
	"time"
)

//channel 实现父线程 等待所有子线程执行完再执行
//todo 程序执行完会报错 fatal error: all goroutines are asleep - deadlock!
func sonv11(done chan bool){

	fmt.Println("sonv11 work...")
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
	}
	fmt.Println("sonv11 work done")
	done <- true
}

func sonv22(done chan bool){

	fmt.Println("sonv22 work...")
	time.Sleep(3*time.Second)
	fmt.Println("sonv22 work done")
	done <- true
}

func main(){
	done := make(chan bool,2)
	defer close(done)

	go sonv11(done)
	go sonv22(done)

	for v := range done {
		fmt.Println(v)
	}

	fmt.Println("all son work done")
}
