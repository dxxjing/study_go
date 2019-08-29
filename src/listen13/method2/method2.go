package main

import (
	"fmt"
)

//继承

type Animal struct{

}

func (a *Animal) Dis(){
	fmt.Println("animal dis")
}

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

func (p *People) Dis(){
	fmt.Println("people dis")
}

type Student_1 struct{
	Score int
	People
}

func (s *Student_1) Set(name string,sex int,score int) {
	s.People.Set(name,sex)//调用父类set
	s.Score = score
}

func test1(){
	s := Student_1{}
	s.Set("jdx",1,98)
	fmt.Println(s)//{98 {jdx 1}}
}

type Student_2 struct{
	Score int
	*People //若不声明为指针，传参是值传递，修改该结构体的成员将无效（因为修改的是副本）
	*Animal
}

func (s *Student_2) Set(name string,sex int,score int) {
	s.People.Set(name,sex)//调用父类set
	s.Score = score
}

func test2(){
	//当结构体中含有匿名成员指针时 尽量将该结构体初始化为指针
	s := &Student_2{}
	//因Student_2中People为匿名成员，所以必须要先分配地址 否则报错
	s.People = &People{}

	s.Set("jdx",1,98)
	fmt.Println(s)//{98 0xc000050460}
	fmt.Println(s.People)//&{jdx 1}

	//Student_2有Show方法调用自己的，没有则会自动调用父类People的Show方法
	s.Show()//{jdx 1}

}

func test3(){
	s := &Student_2{}
	s.People = &People{}
	s.Animal = &Animal{}
	//直接调用会报错 因其父类 people animal 均有Dis方法 底层不知道调用哪个
	//s.Dis()
	//正确的方式
	s.People.Dis()//people dis
	s.Animal.Dis()//animal dis
}

func main(){
	test1()
	test2()
	test3()
}
