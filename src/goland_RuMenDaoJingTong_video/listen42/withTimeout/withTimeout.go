package main

import (
	"context"
	"log"
	"os"
	"time"
)

var lg *log.Logger
//context.withTimout() 返回cancel函数 调用cancel 会触发context.Done()

func do(cx context.Context){
	for{
		time.Sleep(time.Second)
		select{
		case <-cx.Done() :
			lg.Printf("done")
			return //返回主线程 防止重复调用(虽然可以重复调用)
		default:
			lg.Printf("work")
		}
	}
}

func main(){
	lg = log.New(os.Stdout,"",log.Ltime)
	//超时调用cancel  goroutine因超时返回  以下两种 效果一样
	//cx,cancel := context.WithTimeout(context.Background(),5 * time.Second)
	cx,cancel := context.WithDeadline(context.Background(),time.Now().Add(5*time.Second))
	defer cancel()

	go do(cx)

	time.Sleep(10 * time.Second)
}
/*
19:45:40 work
19:45:41 work
19:45:42 work
19:45:43 work
19:45:44 done

 */
