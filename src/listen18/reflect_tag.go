package main

import (
	"fmt"
	"reflect"
)

type Student struct{
	Name string `species:"jdx" color:"red" json:"fff"`
}

//获取结构体tag
func testTag(){
	//传值
	var s Student
	t := reflect.TypeOf(s)
	field,_ := t.FieldByName("Name")
	fmt.Println(field.Tag.Get("species"),field.Tag.Get("color"),field.Tag.Get("json"))

	//传入指针 t2.Elem() 类似 *t2
	var s2 Student
	t2 := reflect.TypeOf(&s2)
	field2,_ := t2.Elem().FieldByName("Name")
	fmt.Println(field2.Tag.Get("species"),field2.Tag.Get("color"),field2.Tag.Get("json"))
}

func main(){
	testTag()
}
