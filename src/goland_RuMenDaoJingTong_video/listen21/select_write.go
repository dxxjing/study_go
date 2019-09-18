package main

import (
	"fmt"
	"time"
)
//当读慢  写快时  channel就会有满的时候
//select 可以用于判断channel是否满
func write(ch chan string){
	for {
		time.Sleep(time.Second)
		select {
		case ch <- "hello":
				fmt.Println("write success")
		default://当channel 满时 执行default
			fmt.Println("channel full")
		}
	}
}

func main(){
	ch := make(chan string,10)

	go write(ch)
	for v := range ch {
		time.Sleep(time.Second *2)
		fmt.Println(v)
	}


}
