package main

import (
	"fmt"
)

//含有基本类型的匿名字段 使用类型作为字段名
//相同类型的匿名字段只能有一个
type User struct {
	ID int
	Name string
	int
	string
}

//基本类型的匿名字段
func structV1(){
	var u User
	u.ID = 1
	u.Name = "jdx"
	u.int = 23
	u.string = "nimingziduan"
	fmt.Println(u) //{1 jdx 23 nimingziduan}

	u2 := User{
		ID : 2,
		Name : "tom",
		int : 33,
		string : "nihao",
	}
	fmt.Println(u2)//{2 tom 33 nihao}
}

type Address struct {
	Province string
	City string
}

//嵌套结构体
type User2 struct {
	ID int
	Name string
	addr Address
}
func user2Test(){
	u := User2{
		ID : 1,
		Name : "jdx",
		addr : Address{ //这里必须要在括号外指定内嵌类型
			Province : "安徽",
			City : "阜阳",
		},
	}
	fmt.Println(u)//{1 jdx {安徽 阜阳}}
	//访问
	id := u.ID
	ad := u.addr.Province
	ct := u.addr.City
	fmt.Println(id,ad,ct)
}

type User3 struct {
	ID int
	Name string
	Address //嵌套匿名结构体
}

func user3Test(){
	u := User3{
		ID : 1,
		Name : "tom",
		Address:Address{ //类型名 作为字段名
			Province:"上海",
			City:"嘉定",
		},
	}
	fmt.Println(u)//{1 tom {上海 嘉定}}
	//匿名嵌套结构体字段的访问
	n := u.Name
	ct := u.City
	pr := u.Province
	fmt.Println(n,ct,pr)//tom 嘉定 上海
}

//嵌套具名结构体指针
type User4 struct {
	ID int
	Name string
	addr *Address
}

func modify4_1(u User4){
	u.Name = "new"
	u.addr.City = "界首"
}
func modify4_2(u *User4){
	u.Name = "new"
	u.addr.City = "界首"
}

func user4Test(){
	u := User4{
		ID : 1,
		Name : "jdx",
		addr : &Address{//当嵌套结构体指针时 必须指定字段名
			Province:"安徽",
			City:"合肥",
		},
	}
	fmt.Println(u)//{1 jdx 0xc000004440} 最后是个地址

	//访问字段
	fmt.Println(u.addr.City,u.addr.Province)//合肥 安徽
	//结构体是值传递 所以name 没有改变
	modify4_1(u)
	fmt.Println(u.addr.City,u.Name)//界首 jdx
	//传入地址 改变name
	modify4_2(&u)
	fmt.Println(u.addr.City,u.Name)//界首 new


}

type User5 struct{
	ID int
	Name string
	*Address
}

func user5Test(){
	var u User5
	u.ID = 1
	u.Name = "jdx"
	//嵌套匿名结构体指针 必须要分配地址

	/*//第一种方式
	u.Address = new(Address)
	u.Province = "上海"
	u.City = "嘉定"
	*/
	//第二种
	u.Address = &Address{
		Province:"安徽",
		City:"合肥",
	}
	fmt.Println(u)
	fmt.Println(u.City,u.Province)
	//{1 jdx 0xc000004440}
	//嘉定 上海

	// 初始化匿名结构体指针必须 以指针的形式初始化
	/*错误方式
	u2 := User5{
			ID:1,
			Name:"jdx",
			Address:&Address{
				Province:"安徽",
				City:"合肥",
			},
		}
	*/
	u2 := &User5{
		ID:1,
		Name:"jdx",
		Address:&Address{
			Province:"安徽",
			City:"合肥",
		},
	}
	fmt.Println(u2)//&{1 jdx 0xc0000504a0}

}

type Stu struct {
	Name string
}
type User6 struct{
	ID int
	Name string
	stu Stu
}

func user6Test(){
	u := User6{
		ID:1,
		Name:"jdx",
		stu:Stu{
			Name:"jdx2",
		},
	}
	fmt.Println(u)//{1 jdx {jdx2}}

	//访问
	fmt.Println(u.Name,u.stu.Name)//jdx jdx2
}

type User7 struct{
	ID int
	Name string
	Stu
}

func user7Test(){
	u := User7{
		ID:1,
		Name:"jdx",
		Stu:Stu{
			Name:"jdx2",
		},
	}
	fmt.Println(u)//{1 jdx {jdx2}}
	//访问
	fmt.Println(u.Name,u.Stu.Name)//jdx jdx2

}

func main(){
	//structV1()
	//user2Test()
	//user3Test()
	//user4Test()
	user5Test()
	//user6Test()
	//user7Test()
}
