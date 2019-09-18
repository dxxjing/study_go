package main

import (
	"fmt"
	"time"
)


func task(ch chan string){
	time.Sleep(time.Second*10)
	ch <- "test2 done"

}

func main(){
	ch := make(chan string)
	go task(ch)

	var str string

	select{
	case str = <- ch:
		fmt.Println(str)
		//超时 当task 5秒没有返回 就会执行超时操作
	case <- time.After(time.Second * 5):
		fmt.Println("time out")
	}
}
