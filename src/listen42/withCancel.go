package main

import (
	"context"
	"log"
	"os"
	"time"
)
var logger *log.Logger
//context.withCancel() 返回cancel函数 调用cancel 会触发context.Done()

func do(cx context.Context){
	for{
		time.Sleep(time.Second)
		select{
		case <-cx.Done(): //收到cancel 发送的channel后 继续执行以下代码
			logger.Printf("done")
		default:
			logger.Printf("work")
		}
	}
}

func main(){
	logger = log.New(os.Stdout,"",log.Ltime)

	cx,cancel := context.WithCancel(context.Background())

	go do(cx)

	//10秒后调用cancel 发送终止信号给goroutine 但不会立即终止 是等待主程序结束goroutine才结束
	time.Sleep(10 * time.Second)
	cancel() //向channel 发送数据

	logger.Printf("down")

	time.Sleep(5 * time.Second)
}
/*
19:30:58 work
19:30:59 work
19:31:00 work
19:31:01 work
19:31:02 work
19:31:03 work
19:31:04 work
19:31:05 work
19:31:06 work
19:31:07 down
19:31:07 done
19:31:08 done
19:31:09 done
19:31:10 done
19:31:11 done

 */
