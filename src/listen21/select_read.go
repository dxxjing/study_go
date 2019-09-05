package main

import (
	"fmt"
	"time"
)

func test(ch chan string){
	time.Sleep(time.Second * 6)
	ch <- "test done"
}

func test2(ch chan string){
	time.Sleep(time.Second*3)
	ch <- "test2 done"

}

//select 同时监听一个或多个channel 直到其中一个channel ready
//监听多个channel 同时ready  随机选择一个执行
//select 在channel 为空时 会阻塞等待
func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)
	go test(ch1)
	go test2(ch2)

	var str string

	select{
	case str = <- ch1:
		fmt.Println(str)
	case str = <- ch2:
		fmt.Println(str)
	}
}
