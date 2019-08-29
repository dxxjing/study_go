package main

import (
	"fmt"
	"os"
)
/*
实现简单的学生管理系统 每个学生有分数 年级 性别 姓名等字段，
用户可以在控制台添加学生、修改学生信息、打印所有学生列表
 */



type Student struct {
	Name string
	Grade string
	Score int
	Sex int
}
//采用map存储
type StuList map[string]Student //key 为 name

func NewStudent(name,grade string,score,sex int) *Student {
	return &Student{
		Name:name,
		Grade:grade,
		Score:score,
		Sex:sex,
	}
}


func (list *StuList) Add(name,grade string,score,sex int){
	stu := NewStudent(name,grade,score,sex)
	(*list)[name] = *stu
}

func (list *StuList) Del(name string){
	delete(*list,name)
}

func (list *StuList) Show(){
	for _,v := range *list {
		fmt.Printf("name:%s,sex:%d,grade:%s,score:%d \n",v.Name,v.Sex,v.Grade,v.Score)
	}
}

func InputInfo()(string,string,int,int){
	var(
		name string
		grade string
		score int
		sex int
	)
	fmt.Printf("请输入姓名:")
	fmt.Scanf("%s\n",&name)
	fmt.Printf("请输入年级[1-6]:")
	fmt.Scanf("%s\n",&grade)
	fmt.Printf("请输入分数[0-100]:")
	fmt.Scanf("%d\n",&score)
	fmt.Printf("请输入性别[0|1]:")
	fmt.Scanf("%d\n",&sex)
	return name,grade,score,sex
}

func ShowMenu(){
	fmt.Printf("请输入：")
	fmt.Printf("1 列表，2 添加,3 修改，4 删除,5 退出 \n")
}

func main(){
	var num int
	var list StuList
	list = make(StuList,100)

	for {
		ShowMenu()
		fmt.Scanf("%d\n",&num)
		switch num {
		case 1:
			list.Show()
		case 2,3:
			name,grade,score,sex := InputInfo()
			list.Add(name,grade,score,sex)
		case 4:
			var name string
			fmt.Printf("请输入姓名:")
			fmt.Scanf("%s\n",&name)
			list.Del(name)
		case 5:
			os.Exit(0) //退出
		default:
			list.Show()
		}
	}
}



