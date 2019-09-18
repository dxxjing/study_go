package main

import (
	"fmt"
)

type User struct {
	Name string
	ID int
	Addr string
}

//结构体中的字段名 不能加引号
//结构体 值或者指针 都用 . 访问成员


//三种初始化方式
func createStruct(){
	var u User = User{
		Name : "jdx",
		ID : 2,
		Addr : "上海",
	}
	fmt.Println(u) //{jdx 2 上海}

	u2 := User{
		Name : "jdx",
		ID : 2,
		Addr : "上海",
	}
	fmt.Println(u2)

	var u3 User
	u3.Name = "tom"
	u3.ID = 12
	u3.Addr = "安徽"
	fmt.Println(u3)
	//取值
	fmt.Println(u3.Name)

}

//结构体指针
func structPoint(){
	var u *User
	u = &User{
		ID :14,
		Name : "jdx",
		Addr : "山东",
	}
	fmt.Println(u)//&{jdx 14 山东} //结果前有取地址符 表示该变量为指针

	u2 := &User{
		ID :14,
		Name : "jdx",
		Addr : "山东",
	}
	fmt.Println(u2)

	//声明指针 必须要初始化才能操作 否则panic
	var u3 *User
	u3 = new(User) //初始化
	u3.Name = "baby"
	u3.ID = 19
	u3.Addr = "河南"
	fmt.Println(u3)//&{baby 19 河南}

	//结构体指针 (*u4).Name  和 u4.Name 都可使用
	var u4 = new(User)
	(*u4).Name = "lala"
	u4.ID = 18
	(*u4).Addr = "福建"
	fmt.Println(u4)//&{lala 18 福建}
	//取值
	fmt.Println(u4.Name)//lala
}

//构造函数
//一般都是返回指针 节省空间 传递效率高
func NewUser(name string,id int,addr string) *User {
	return &User{
		Name : name,
		ID : id,
		Addr : addr,
	}
}

func NewUserV2(name string,id int,addr string) *User {
	var u *User
	u = new(User)
	u.Name = name
	u.ID = id
	u.Addr = addr
	return u
}

func main(){
	//createStruct()
	//structPoint()
	u1 := NewUser("jdx",18,"安徽")
	u2 := NewUserV2("tom",12,"上海")
	fmt.Println(u1,u2) //&{jdx 18 安徽} &{tom 12 上海}
}
