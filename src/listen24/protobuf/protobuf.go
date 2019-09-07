package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"listen24/protobuf/address"
)

/*
	window下安装
	https://github.com/protocolbuffers/protobuf/releases
	1.下载 window 对应版本，解压 将bin目录下protoc.exe 拷贝到gopath bin目录下(E:\study_go\bin)
	2.将E:\study_go\bin 添加到系统环境变量Path
	3 go get -u github.com/golang/protobuf/protoc-gen-go
	命令生成文件
	cd src/listen24/protobuf
	protoc --go_out =./address/ ./person.proto
 */

func main(){
	/*
	err := pbMarshal()
	if err != nil {
		fmt.Println(err)
		return
	}

	 */
	data,err := pbUnmarshal()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

func pbUnmarshal() (*address.ContactBook,error) {
	var contact address.ContactBook

	pbData,err := ioutil.ReadFile("./pb.txt")
	if err != nil {
		return nil,err
	}

	err = proto.Unmarshal(pbData,&contact)
	if err != nil {
		return nil,err
	}
	fmt.Println(contact)
	return &contact,nil
}

func pbMarshal() error {
	var contactbook address.ContactBook
	for i := 0; i < 20; i++ {

		p := address.Person{
			Id:int32(i),
			Name:fmt.Sprintf("jing-%s",i),
		}

		ph := address.Phone{
			Type:address.PhoneType_HOME,
			Number:fmt.Sprintf("1770210523%d",i),
		}
		p.Phones = append(p.Phones,&ph)

		contactbook.Persons = append(contactbook.Persons,&p)
	}
	//生成二进制数据
	pbData,err := proto.Marshal(&contactbook)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./pb.txt",pbData,0755)

	return err
}

