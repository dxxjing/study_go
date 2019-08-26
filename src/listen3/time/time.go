package main

import (
	"fmt"
	"time"
)

func work(){
	fmt.Println("task work:")
}

//该定时任务 无法关闭
func tickTask(){
	ch := time.Tick(3 * time.Second) //3 秒定时任务  return channel time
	for k := range ch {
		fmt.Printf("%v \n",k)
		work()
	}
}

func tickTaskV2(){
	ticker := time.NewTicker(3 * time.Second)
	i := 0 //五次后关闭
	for k := range ticker.C {
		i += 1
		fmt.Printf("v2: %v \n",k)
		work()
		if i == 5 {
			ticker.Stop()
		}
	}
}

func main(){

	//tickTask()
	//tickTaskV2()

	dur,_ := time.ParseDuration("-1.5h")
	fmt.Printf("-1.5h : %v \n",dur)
}