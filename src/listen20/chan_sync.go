package main

import (
	"fmt"
	"time"
)

//使用channel 完成父子线程同步
func son(ch chan bool){
	//模拟耗时的任务
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
	}
	//任务执行完 向channel 写入 true
	ch <- true
}

func main(){
	//创建阻塞channel  读 无数据 阻塞等待, 写 已满 也阻塞等待
	ch := make(chan bool)
	defer close(ch)
	//子协程
	go son(ch)

	<- ch //阻塞等待子协程执行完

	fmt.Println("son done")

}
