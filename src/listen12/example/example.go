package main

import (
	"flag"
	"fmt"
)
/*
实现简单的学生管理系统 每个学生有分数 年级 性别 姓名等字段，
用户可以在控制台添加学生、修改学生信息、打印所有学生列表
 */

var (
	action string
	pName string
	pGrade string
	pScore int
	pSex int
)


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

func (list *StuList) Edit(name,grade string,score,sex int){
	if v,ok := (*list)[name];ok {
		if grade != "" {
			v.Grade = grade
		}
		if score != -1 {
			v.Score = score
		}
		if sex != -1 {
			v.Sex = sex
		}
		(*list)[name] = v
	}else{
		fmt.Println("the student not found")
	}

}

func (list *StuList) Del(name string){
	delete(*list,name)
}

func (list StuList) Show(){
	for _,v := range list {
		fmt.Printf("name:%s,sex:%d,grade:%s,score:%d \n",v.Name,v.Sex,v.Grade,v.Score)
	}
}

func ParseFlag(){
	flag.StringVar(&action,"a","show","add 添加 edit 修改 del 删除 show 展示 exit 退出")
	flag.StringVar(&pName,"name","","姓名")
	flag.StringVar(&pGrade,"grade","","年级")
	flag.IntVar(&pScore,"score",-1,"分数")
	flag.IntVar(&pSex,"sex",-1,"性别")
	flag.Parse()
}

func main(){
	ParseFlag()

	var list StuList
	list = make(StuList,100)
	for {
		switch action {
		case "show":
			list.Show()
		case "add":
			list.Add(pName,pGrade,pScore,pSex)
			list.Show()
		case "edit":
			list.Edit(pName,pGrade,pScore,pSex)
			list.Show()
		case "del":
			list.Del(pName)
			list.Show()
		case "exit":
			break
		default:
			list.Show()
		}
	}
}



