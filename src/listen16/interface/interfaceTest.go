package main

import (
	"fmt"
)

//定义一个Animal 接口
type Animal_v1 interface {
	Eat()
	Talk()
	Name() string
}

type Animal_v2 interface {
	Run()
}

type Animal interface {
	Animal_v2
	Animal_v1
}

//定义Dog结构体 让其实现Animal 接口
type Dog struct{

}
func (d Dog) Eat(){
	fmt.Printf("吃骨头\n")
}
func (d Dog) Talk(){
	fmt.Printf("汪汪汪\n")
}
func (d Dog) Name() string{
	n := "阿黄"
	fmt.Printf("%s\n",n)
	return n
}
func (d Dog) Run(){
	fmt.Printf("dog run\n")
}

//定义Pig结构体 并实现Animal接口
type Pig struct{

}
func (d Pig) Eat(){
	fmt.Printf("吃猪草\n")
}
func (d Pig) Talk(){
	fmt.Printf("哼哼哼\n")
}
func (d Pig) Name() string{
	n := "佩奇"
	fmt.Printf("%s\n",n)
	return n
}

func test_interface(){
	var d Dog
	var a Animal_v1
	a = d

	a.Eat()
	a.Talk()
	a.Name()

	var p Pig
	a = p
	a.Eat()
	a.Talk()
	a.Name()
}

func test_interface2(){
	var a Animal_v1
	//指针类型实现接口 和 值类型实现接口 的区别

	//1.值类型实现接口，可以将 值 或 指针 赋值给接口变量
	var d2 Dog
	a = d2
	a.Eat()

	var d3 *Dog
	d3 = &Dog{} //指针必须初始化
	a = d3
	a.Eat()
	//2.指针类型实现接口 只能将指针赋值给接口变量
	//原因：一个变量存储在接口变量中就不能获取该变量的地址
	//验证 ：将Dog 的所有方法接收者更改为指针类型
}

func test_v1(){
	//一个类型 可以实现多个接口 即若Dog类型实现Animal_v1和Animal_v2接口 则Dog可以赋值给Animal_v1或Animal_v2
	var a1 Animal_v1
	var a2 Animal_v2
	var d Dog
	a1 = d
	a1.Eat()
	a2 = d
	a2.Run()
}

func testv2(){
	//接口嵌套 Animal接口 内有 Animal_v1 Animal_v2
	//则若Dog 要想实现Animal接口 则必须同时实现 Animal_v1 Animal_v2 接口
	var a Animal
	var d Dog
	a = d
	a.Eat()
	//Pig 没有实现 Animal接口中的Run 不能赋值
	//var p Pig
	//a = p
}

func main(){
	//test_interface()
	//test_interface2()
	//test_v1()



}


