package main

import (
	"encoding/json"
	"fmt"
)
//结构体序列化 json

type Student struct {
	ID string
	Name string
	Score int
}

type Class struct {
	Name string
	Count int
	Stu []*Student
	//test2 test3 证明[]Student，[]*Student 没有太大区别，都可直接修改
	//切片本身就是传引用 直接传切片 和 指针没有区别
}

func test1(){
	c := Class{
		Name:"一年级",
		Count:10,
	}
	//添加10个学生
	for i := 0; i < 10; i++ {
		s := &Student{
			ID : fmt.Sprintf("%d",i),
			Name : fmt.Sprintf("stu_%d",i),
			Score : 90 + i,
		}
		c.Stu = append(c.Stu,s)
	}
	//json序列化
	data,err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal err")
	}
	fmt.Printf("json:%s\n",string(data))
	/*
		{"Name":"一年级","Count":10,"Stu":[{"ID":"0","Name":"stu_0","Score":90},{"ID":"1","Name":"stu_1","Score":91},{"ID":"2","Name"
		:"stu_2","Score":92},{"ID":"3","Name":"stu_3","Score":93},{"ID":"4","Name":"stu_4","Score":94},{"ID":"5","Name":"stu_5","Score":95
		},{"ID":"6","Name":"stu_6","Score":96},{"ID":"7","Name":"stu_7","Score":97},{"ID":"8","Name":"stu_8","Score":98},{"ID":"9","Name":
		"stu_9","Score":99}]}
	*/
	//反序列化
	rowJson := `{"Name":"一年级","Count":10,"Stu":[{"ID":"0","Name":"stu_0","Score":90},{"ID":"1","Name":"stu_1","Score":91},{"ID":"2","Name"
	:"stu_2","Score":92},{"ID":"3","Name":"stu_3","Score":93},{"ID":"4","Name":"stu_4","Score":94},{"ID":"5","Name":"stu_5","Score":95
	},{"ID":"6","Name":"stu_6","Score":96},{"ID":"7","Name":"stu_7","Score":97},{"ID":"8","Name":"stu_8","Score":98},{"ID":"9","Name":
	"stu_9","Score":99}]}`
	c2 := &Class{}
	//第二个参数 传入地址 结果会保存在c2
	err = json.Unmarshal([]byte(rowJson),c2)
	if err != nil {
		fmt.Println("json unmarshal err")
	}
	fmt.Println("unmarshal:")
	for _,v := range c2.Stu {
		fmt.Printf("%v",*v)
	}
}

type ClassV2 struct {
	Name string
	Count int
	Stu []Student
}

func (c *ClassV2) Modify(){
	c.Stu[0].Name = "jdx"
}

type ClassV3 struct {
	Name string
	Count int
	Stu []*Student
}

func (c *ClassV3) Modify(){
	c.Stu[0].Name = "jdx"
}

func test2(){
	c := &ClassV2{
		Name:"class1",
		Count:20,
	}
	c.Stu = append(c.Stu,Student{ID:"123",Name:"tom",Score:89})
	c.Modify()
	fmt.Println(c.Stu)//[{123 jdx 89}]

}

func test3(){
	c := &ClassV3{
		Name:"class1",
		Count:20,
	}
	c.Stu = append(c.Stu,&Student{ID:"123",Name:"tom",Score:89})
	c.Modify()
	fmt.Println(c.Stu[0])//&{123 jdx 89}

}

func main(){
	//test1()
	test2()
	test3()
}
