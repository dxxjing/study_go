package main

import (
	"fmt"
	"time"
)

//获取代码执行时间 单位 us 微秒
func testCost(){
	start := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()
	cost := (end - start) / 1000 //纳秒转为微秒
	fmt.Printf("cost %v us \n",cost)
}

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

	testCost()

	//特殊点 格式化时间
	curTime := time.Now()
	//这里的年月日时分秒 必须是以下的时间 否则会有问题
	//2006-01-02 15:04:05 是go语言诞生的时间点
	timeFormat := curTime.Format("2006-01-02 15:04:05") 
	fmt.Printf("format:%v \n",timeFormat) //2019-08-26 12:48:13
	//修改任何一个都会返回格式化错误的时间
	timeFormat = curTime.Format("2009-01-02 15:04:05") 
	fmt.Printf("format:%v \n",timeFormat) //26009-08-26 12:48:13  错误

}