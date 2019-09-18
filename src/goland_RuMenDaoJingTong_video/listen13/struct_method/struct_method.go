package main

import (
	"fmt"
)

//指针接受者 和值接收者

type People struct {
	Name string
	Sex int
}
//指针接收者
func (p *People) Set(name string,sex int){
	p.Name = name
	p.Sex = sex
}

//值接收者
func (p People) Show(){
	fmt.Println(p)
}

func test1(){
	p := People{}
	//go 底层自动将p.Set() 转为 (&p).Set()
	p.Set("jdx",1)
	fmt.Println(p)//{jdx 1}
	(&p).Set("tom",2)//{tom,2}
}

func test2(){
	p := &People{}
	p.Set("jdx",1)

	//go底层自动将p.Show() -> (*p).Show()
	p.Show()//{jdx 1}
	(*p).Show()//{jdx 1}
}

func main(){
	//test1()
	test2()

}
