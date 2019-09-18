package main

import (
	"fmt"
)

func add(a,b int) int {
	return a + b
}

func sub(a,b int) int {
	return a - b
}

//函数作为参数 
func test(a,b int,op func(int,int)int) int {
	return op(a,b)
}

func main(){
	//通过传入不同的函数 进行不同的操作
	s := test(1,2,add)
	d := test(3,1,sub)
	fmt.Printf("s = %d, d = %d \n",s,d)
}