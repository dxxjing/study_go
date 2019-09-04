package main

import (
	"fmt"
	"sync"
	"time"
)

//父线程 等待所有子线程执行完再执行

func sonv1(wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Println("sonv1 work...")
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
	}
	fmt.Println("sonv1 work done")

}

func sonv2(wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Println("sonv2 work...")
	time.Sleep(3*time.Second)
	fmt.Println("sonv2 work done")
}

func main(){
	var wg sync.WaitGroup
	//Add 方法必须在父线程中执行
	wg.Add(2)
	//必须要传入地址进去 否则执行完会报错fatal error: all goroutines are asleep - deadlock!
	go sonv1(&wg)
	go sonv2(&wg)

	wg.Wait()
	fmt.Println("all son work done")
	
}
