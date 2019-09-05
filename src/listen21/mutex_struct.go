package main

import(
	"fmt"
	"sync"
	"time"
)
//互斥锁

//***************************************************
//todo 注意 锁为结构体，属于值传递，需要加锁 必须使用指针作为接受者 否则加锁无效 *
//***************************************************

//结论:map是引用传递 所以结构体中用 值 和 指针 都一样
type User struct{
	//Score map[string]int //值map
	Score *map[string]int  //指针map
	mu sync.Mutex
}

func (u *User) Edit(key string,num int){
	u.mu.Lock()
	fmt.Printf("%dth lock\n",num)
	defer u.mu.Unlock()

	time.Sleep(time.Second)//模拟耗时较长
	//u.Score[key]++    //值map ++
	(*(u.Score))[key]++ //指针map ++
	fmt.Printf("%dth unlock\n",num)
}

func main(){
	u := User{}
	//u.Score = make(map[string]int)
	score := make(map[string]int)//这里必须赋值给变量才能取地址
	u.Score = &score
	//启用五个协程并发 修改结构体
	for i := 0; i < 5; i++ {
		go u.Edit("math",i)
	}
	//这里保证主程序 晚于 子协程结束
	time.Sleep(time.Second * 6)
	//fmt.Println(u.Score)
	fmt.Println(*(u.Score))
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

