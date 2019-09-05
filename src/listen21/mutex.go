package main

import (
	"fmt"
	"sync"
)

//多线程 并发访问同一份资源 就要采取同步措施 比如使用互斥锁
//锁 和 waitgroup 都是结构体 所以必须传地址
func syncTest(num *int,wg *sync.WaitGroup,mu *sync.Mutex){
	for i := 0; i < 5000; i++{
		mu.Lock()
		*num = *num +1
		mu.Unlock()
	}
	wg.Done()
}

func unSyncTest(num *int,wg *sync.WaitGroup){
	for i := 0; i < 5000; i++{
		*num = *num +1
	}
	wg.Done()
}

//不加锁 每次执行后的num 都不一样
//加锁后 每次执行 结果都一样 且都等于20000
func main(){
	var mutex sync.Mutex
	var num int
	var wg sync.WaitGroup
	wg.Add(4)
	//启用四个协程 并发执行

	go syncTest(&num,&wg,&mutex)
	go syncTest(&num,&wg,&mutex)
	go syncTest(&num,&wg,&mutex)
	go syncTest(&num,&wg,&mutex)
	/*
	go unSyncTest(&num,&wg)
	go unSyncTest(&num,&wg)
	go unSyncTest(&num,&wg)
	go unSyncTest(&num,&wg)
	*/
	wg.Wait()

	fmt.Println(num)


}
