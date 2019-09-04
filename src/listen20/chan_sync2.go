package main

import (
	"fmt"
	"time"
)

//channel 实现父线程 等待所有子线程执行完再执行
//todo 结论:channel 必须在写端进行关闭，因为只有写端才知道何时不用

//两个协程 同时向同一个管道写数据，谁负责关闭？
//无法判断谁后结束 所以尽量使用waitGroup 进行同步
func sonv11(done chan bool){

	fmt.Println("sonv11 work...")
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
	}
	fmt.Println("sonv11 work done")
	done <- true
	close(done)
}

func sonv22(done chan bool){

	fmt.Println("sonv22 work...")
	time.Sleep(3*time.Second)
	fmt.Println("sonv22 work done")
	done <- true
	//close(done)
}

func main(){
	done := make(chan bool,2)
	//defer close(done)

	go sonv11(done)
	go sonv22(done)

	//todo for range完毕 程序执行完会报错 fatal error: all goroutines are asleep - deadlock!
	/*初步预测：由于主程序不知道何时停止从管道读取数据 且读取两次后 所有子协程都已结束，
		主程序仍在等待从管道读取数据 导致死锁deadlock
	*****解决办法：使用固定次数的for循环
	第二种方法：
		在写入的程序中关闭,使用for range自动判断channel关闭，终止循环，就不会出现死锁,
	*/
	for v := range done {
		fmt.Println(v)
	}

	//解决办法1
	/*for i := 0; i < 2; i++ {
		fmt.Println(<-done)
	}*/

	//方法3 该方法在写端不关闭channel的情况下 仍然会出现死锁
	/*for{
		v,ok := <- done
		if !ok {
			break
		}
		fmt.Println(v)
	}*/

	fmt.Println("all son work done")
}
