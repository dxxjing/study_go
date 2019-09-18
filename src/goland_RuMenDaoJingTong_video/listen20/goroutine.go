package main

import (
	"fmt"
	"time"
)

func main(){
	//goroutine 协程 并发执行
	//谁先执行看内部协程调度 谁先抢到cpu 谁先执行 所以一下输出 没有规律
	for i := 0; i < 10; i++ {
		go fmt.Printf("%d ",i)
	}
	fmt.Println("this is main pthread")
	//主线程sleep原因：主线程结束，goroutine会被强制退出，无法执行
	time.Sleep(time.Second)
}
