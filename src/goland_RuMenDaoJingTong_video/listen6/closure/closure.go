package main 

import (
	"fmt"
	"strings"
	"time"
)

//闭包 该函数返回匿名函数 
func closure_v1(num int) func(int) int {
	return func (x int) int {
		num += x
		return num
	}
}

//闭包应用
func closure_v2(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name,suffix) {
			name = name + suffix
		}
		return name
	}
}
//返回 两个匿名函数
func closure_v3(num int) (func(int) int,func(int) int){
	add := func(i int) int {
		num += i
		return num
	}

	sub := func(i int) int {
		num -= i
		return num
	}

	return add,sub
}

func closure_v4(){
	//坑点：i 应通过匿名函数传入 否则 打印的全是5
	for i := 0; i < 5; i++ {
		go func(i int){
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}

func main(){
	// 闭包中的变量 num 生命周期和变量f1一样，num就像f1中的成员变量一样 
	f1 := closure_v1(5)// num : 5
	s1 := f1(1) //num : 6 = 5 + 1
	s2 := f1(2) //num : 8 = 6 + 2
	fmt.Printf("closure s1 = %d,s2 = %d \n",s1,s2) 

	f2 := closure_v2(".jpg")
	f3 := closure_v2(".php")
	fmt.Printf("closure v2 : %s , %s \n",f2("test"),f3("test")) // test.jpg test.php

	closure_v4()
}
