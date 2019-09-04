package main

import (
	"fmt"
	"time"
)
//todo 注意 关闭channel 一定要在*写入*的程序中关闭，否则就会出现死锁 deadlock
//如该程序 必须在test函数中关闭,若在main中关闭 就会出现死锁
func test(ch chan int){
	defer close(ch)
	for i := 0; i < 10; i++ {
		time.Sleep(2*time.Second)
		ch <- i
	}
}

func main(){
	ch := make(chan int,10)
	//defer close(ch)

	go test(ch)

	for v := range ch{
		fmt.Println(v)
	}
	fmt.Println("test done")
}
