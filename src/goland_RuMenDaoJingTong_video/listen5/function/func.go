package main

import (
	"fmt"
)

//可变参数 可传入 大于等于0个参数
func calc_v1(a ...int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

//可变参数 可传入 大于等于1个参数 b 为固定参数 必须带类型
func calc_v2(b int,a ...int) int {
	sum := b
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func testDefer(){
	a := 0
	defer fmt.Printf("defer a = %d \n",a) // 0
	//匿名函数 当调用匿名函数时a已经是100
	defer func(){
		fmt.Printf("defer func a = %d \n",a) // 100
	}()

	a = 100

}

func main(){
	//s := calc_v1()
	s := calc_v2(1,2,3)
	fmt.Printf("calc s = %d \n",s)

	testDefer()
}