package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

//结构体转 json  所有成员必须是可导出的
type Rsp struct {
	Code int	`json:"code"`
	Msg	string	`json:"msg"`
	Data []*User`json:"data"`
}

type User struct{
	Name string	`json:"name"`
	Age int		`json:"age"`
	Sex string	`json:"sex"`
}
/*
	{
		"code":0,
		"msg":"success",
		"data":[
			{
				"Name":"name0",
				"Age":81,
				"Sex":"man"
			},
			{
				"Name":"name1",
				"Age":87,
				"Sex":"man"
			}
		]
	}
 */
func writeJson(filename string) (err error) {
	users := genRsp()

	jsonb,err := json.Marshal(*users)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, jsonb,0755)
	if err != nil{
		return err
	}

	return
}
//返回值中定义的变量 如果是指针 需要进行初始化 否则无法取其地址
//也就是说 返回值中定义的指针 rsp 和 rsp = &Rsp{} 是不同的
//为避免踩坑 返回值中不要定义变量
func readJson2(filename string) (rsp *Rsp,err error) {
	rsp = &Rsp{} //必不可少

	jsonb,err := ioutil.ReadFile(filename)
	if err != nil {
		return nil,err
	}
	//第二个参数必须是指针 传出参数
	err = json.Unmarshal(jsonb,rsp)
	return rsp,err
}

func readJson(filename string) (*Rsp,error) {
	var rsp Rsp
	jsonb,err := ioutil.ReadFile(filename)
	if err != nil {
		return nil,err
	}
	//第二个参数必须是指针 传出参数
	err = json.Unmarshal(jsonb,&rsp)
	return &rsp,err
}

func main(){

	err := writeJson("./test.txt")
	if err != nil {
		fmt.Println("write json err:",err)
		return
	}
	data,err := readJson("./test.txt")
	if err != nil {
		fmt.Println("read json err:",err)
		return
	}
	for k,v := range data.Data{
		fmt.Println(k,v)
	}

}

func genRsp()(rsp *Rsp){

	rsp = &Rsp{
		Code:0,
		Msg:"success",
	}

	for i :=0; i < 10; i++ {
		u := &User{
			Name:fmt.Sprintf("name%d",i),
			Age:rand.Intn(100),
			Sex:"man",
		}
		rsp.Data = append(rsp.Data,u)
	}
	return
}
