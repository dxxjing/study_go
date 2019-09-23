package main

import "fmt"

var ch = make(chan bool)
var msg string

func f(){
	msg = "hello go"
	ch <- true
	close(ch)
}

func main() {
	go f()
	val := <- ch
	//另一个协程关闭channel 当前协程也能从channel中正确读取数据
	fmt.Println(val,msg)//true hello go
}
