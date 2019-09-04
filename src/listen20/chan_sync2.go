package main

import (
	"fmt"
	"time"
)

//channel 实现父线程 等待所有子线程执行完再执行

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

	//todo for range完毕 程序执行完会报错 fatal error: all goroutines are asleep - deadlock!
	/*初步预测：由于主程序不知道何时停止从管道读取数据 且读取两次后 所有子协程都已结束，
		主程序仍在等待从管道读取数据 导致死锁
	*****解决办法：使用固定次数的for循环
	for v := range done {
		fmt.Println(v)
	}
	*/
	//解决办法
	for i := 0; i < 2; i++ {
		fmt.Println(<-done)
	}
	fmt.Println("all son work done")
}
