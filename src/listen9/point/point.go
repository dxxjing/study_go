package main

import (
	"fmt"
)

func main(){
	//new make 区别
	var p *int
	p = new(int)
	fmt.Printf("new int :addr %p , %v\n",p,*p) //new :addr 0xc00000e0b0 , 0

	var s *[]int //定义一个切片指针

	s = new([]int)
	fmt.Printf("new slice :addr %p , %v\n",s,*s)//new slice :addr 0xc00005a420 , []
	//new 以后返回值未空的指针 仍需要make 初始化 否则直接访问(*s)[0] 会提示越界
	*s = make([]int,5,10)
	(*s)[1] = 100
	(*s)[2] = 1000

	fmt.Println((*s)[1])




}

