package main

import (
	"fmt"
	"time"
)


func test(){
	var i int;
	for i = 1; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(i)
	}
}

func main(){
	go test()

	time.Sleep(time.Second * 11)
	fmt.Println("main---")
}