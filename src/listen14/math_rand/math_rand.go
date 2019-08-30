package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	由 math.rand包 了解go 标准库文档

	该包包含：
		接口 type Source interface
		结构体 type Rand struct , type Zipf struct 以及一些 结构体方法 包方法

	该包中的结构体及其方法 均可直接使用 无需重新实现 参见 test1

 */

func test1(){
	//直接使用该包中的公共方法、Rand结构体 以及结构体对应的方法 产生随机数
	//time 包： func Now() Time , func (t Time) UnixNano() int64
	//math/rand包： func Seed(seed int64)
	rand.Seed(time.Now().UnixNano()) //产生随机数 如不提前调用Seed() 后面产生的随机数将不变

	n := rand.Int() //产生随机的非负int值
	fmt.Println("rand Int(): ",n)
	n2 := rand.Intn(10) // 产生 大于等于0 且 小于 10 的随机数
	fmt.Println("rand Intn(): ",n2)

}

/*
	再看下time包
	该包中的所有 结构体 自定义类型 常量 包方法 结构体方法 均可直接使用 无需重复实现 参加test2
 */
func test2(){
	//使用常量
	fmt.Println(time.ANSIC) //Mon Jan _2 15:04:05 2006
	fmt.Println(time.Sunday)//Sunday
	//使用包中的方法
	t := time.Now() //返回当前本地时间
	fmt.Println(t)
	//使用Time 结构体 返回时间戳
	fmt.Println(time.Now().Unix())//返回时间戳 秒数  1567148596
	fmt.Println(time.Now().UnixNano()) //返回纳秒数  1567148596492363200

	fmt.Println(time.Now().Zone()) //CST 28800 返回该时区的规范名（如"CET"）和该时区相对于UTC的时间偏移量（单位秒）

	fmt.Println(time.Now().UTC().Unix())
	y,m,d := time.Now().Date() //返回年月日
	fmt.Println(y,m,d)

	//time包中的定时器
	t2 := time.NewTimer(time.Second * 3) //到达3秒向管道Timer.C中发送当前时间
	fmt.Println("newTimer:",<-t2.C) //取出管道中的数据
	//周期性的定时器
	t3 := time.NewTicker(time.Second * 2) //每到2秒 就会像管道发送当前时间
	//遍历管道 取出当前时间
	for v := range t3.C {
		fmt.Println("timer ticker:",v)
	}

}

func main(){
	test1()
	//test2()
}

