package main

import (
	"fmt"
	"reflect"
)

type User struct{
	ID int
	Name string
	Sex int
	//addr string
}

func (u User) T1(){
	fmt.Println("user method t1")
}
func (u User) T2(str string){
	fmt.Println("user method t2:params :",str)
}



func test(a interface{}){
	//法1
	v := reflect.ValueOf(a)
	t := v.Type()
	//法2
	//t := reflect.TypeOf(a)
	k := t.Kind()
	switch k {
	case reflect.Int:
		fmt.Println("a is int")
	case reflect.Struct:
		fmt.Println("a is struct")
		fmt.Println(v.NumField()) //获取结构体中字段数
		//获取结构体中字段名 类型 以及值
		for i := 0; i < v.NumField(); i++{
			//v.Field(i).Interface() == v.Field(i) 获取结构体相关字段的值
			fmt.Println(t.Field(i).Name,v.Field(i).Type(),v.Field(i).Interface(),v.Field(i))
		}
		//ID int 1
		//Name string jdx
		//Sex int 2
	default:
		fmt.Println("unknow")
	}
}

func test2(){
	//改变结构体内字段对应的值
	var u User
	v := reflect.ValueOf(&u) //这里一定要传入地址

	v.Elem().Field(0).SetInt(123)
	v.Elem().FieldByName("Name").SetString("tom")
	v.Elem().FieldByName("Sex").SetInt(2)
	fmt.Println(u)//{123 tom 2}

}

//获取结构体方法 必须是结构体 否则会报错
func test3(){
	var u User
	v := reflect.ValueOf(u) //如果是指针接受者的方法 必须传入地址
	t := v.Type()
	fmt.Println("user method num is %d ",t.NumMethod()) //只统计public方法数量
	//v.Elem().FieldByName("ID").SetInt(222)
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name,t.Method(i).Type)
		//T1 func(main.User)
		//T2 func(main.User)
	}
}

//通过反射调用结构体中的方法
func test4(){
	var u User
	v := reflect.ValueOf(u) //使用指针接收者的方法必须传入地址
	//无参数
	m := v.MethodByName("T1")
	var args []reflect.Value
	m.Call(args)
	//有参数
	m2 := v.MethodByName("T2")
	var args2 []reflect.Value
	param := "jdx"
	paramValue := reflect.ValueOf(param) //转为reflect.Value 类型
	args2 = append(args2,paramValue)
	m2.Call(args2)
}

func main(){
	/*
	u := User{

		ID:1,
		Name:"jdx",
		Sex:2,
	}
	test(u)
	*/
	//test2()
	//test3()
	test4()
}
