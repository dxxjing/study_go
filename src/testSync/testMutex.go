package main

import(
	"fmt"
	"sync"
	"time"
)
//互斥锁

//***************************************************
//todo 注意 需要加锁 必须使用指针作为接受者 否则加锁无效 *
//***************************************************


type User struct{
	Score map[string]int
	mu sync.Mutex
}

func (u *User) Edit(key string,num int){
	u.mu.Lock()
	fmt.Printf("%dth lock\n",num)
	defer u.mu.Unlock()

	time.Sleep(time.Second)//模拟耗时较长

	u.Score[key]++
	fmt.Printf("%dth unlock\n",num)
}

func main(){
	u := User{
		Score:make(map[string]int),
	}
	//u.Score["math"] = 0

	for i := 0; i < 5; i++ {
		go u.Edit("math",i)
	}
	//这里保证主程序 晚于 子协程结束
	time.Sleep(time.Second * 6)
}
/*
加锁后可见:对应序号的lock unlock 都是成对出现
2th lock
2th unlock
4th lock
4th unlock
3th lock
3th unlock
1th lock
1th unlock
0th lock
0th unlock

 */

