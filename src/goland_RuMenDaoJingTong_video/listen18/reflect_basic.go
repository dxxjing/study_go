package main

import(
	"fmt"
	"reflect"
)

func reflect_typeOf(a interface{}){
	t := reflect.TypeOf(a) //类型
	//fmt.Println(t) //int

	//利用反射 判断类型
	v := t.Kind()
	switch v {
	case reflect.Int:
		fmt.Printf("a is int,value = %v\n",a)
	case reflect.Float32:
		fmt.Printf("a is float32,value = %v\n",a)
	default:
		fmt.Printf("unknow\n")
	}

}

func reflect_valueOf(a interface{}){
	v := reflect.ValueOf(a) //值
	t := v.Kind() //类型
	switch t {
	case reflect.Int:
		v.SetInt(16)
		fmt.Printf("a is int,%v %v\n",v,v.Int())
	case reflect.Float32:
		v.SetFloat(3.9)
		fmt.Printf("a is float32,%v %v\n",v,v.Float())
	case reflect.Ptr:
		fmt.Printf("a is pointer\n")
		//通过反射设置值
		v.Elem().SetInt(6)
	default:
		fmt.Printf("unknow\n")
	}
}

func main(){
	var a int
	a = 3
	//reflect_typeOf(a)
	reflect_valueOf(&a)//通过反射设置值 这里要传入地址
	fmt.Println(a)
}
