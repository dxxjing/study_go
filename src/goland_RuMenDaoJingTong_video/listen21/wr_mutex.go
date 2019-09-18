package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)
//读写文件 go内部已使用读写锁 本例不加锁也不会出现 多个协程并发写 文件内容有交叉的情况
//读写锁应用

func read_v1(i int,rwmutex *sync.RWMutex,wg *sync.WaitGroup){

	rwmutex.RLock()
	file,err := os.OpenFile("./test.txt",os.O_CREATE|os.O_RDONLY,0777)
	if err != nil {
		fmt.Println("open file err")
		return
	}
	defer file.Close()

	time.Sleep(time.Second * time.Duration(3 * i))

	var data []byte
	data = make([]byte,128)
	n,err := file.Read(data)
	if err != nil {
		fmt.Printf("%d read err\n",i)
		return
	}
	fmt.Printf("%d read %d byte :%s\n",i,n,string(data))

	rwmutex.RUnlock()
	wg.Done()
}

func write_v1(i int,rwmutex *sync.RWMutex,wg *sync.WaitGroup){

	rwmutex.Lock()
	file,err := os.OpenFile("./test.txt",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0777)
	if err != nil {
		fmt.Println("open file err")
		return
	}
	defer file.Close()
	_,err = file.WriteString(fmt.Sprintf("%d write file",i))
	if err != nil {
		fmt.Printf("%d write err\n",i)
	}
	fmt.Println(i," write success")
	rwmutex.Unlock()
	wg.Done()
}

func main(){

	var wg sync.WaitGroup
	var rwmtx sync.RWMutex


	//启用5协程 并发写
	var j int
	for ;  j< 5; j++ {
		wg.Add(1)
		go write_v1(j,&rwmtx,&wg)
	}

	//启用5个协程 并发读
	var i int
	for ; i < 5; i++ {
		wg.Add(1)
		go read_v1(i,&rwmtx,&wg)
	}


	wg.Wait()
}
