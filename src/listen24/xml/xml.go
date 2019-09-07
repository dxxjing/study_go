package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//结构体所有必须是可导出的
type Config struct{
	//todo 这里必须定义 XMLName 否则执行Marshal 得到的结果和原本的结构不一致
	XMLName xml.Name	`xml:"servers"`
	Version string	`xml:"version,attr"`
	Srv []Server`xml:"server"`
}

type Server struct{
	ServerName string	`xml:"serverName"`
	ServerIP string		`xml:"serverIP"`
}

func wrieXml(){
	conf := &Config{
		Version:"1.0",
	}
	//conf.Srv = append(conf.Srv, Server{"Shanghai_VPN", "127.0.0.1"})
	//conf.Srv = append(conf.Srv, Server{"Beijing_VPN", "127.0.0.2"})
	conf.Srv = []Server{
		Server{
			ServerIP:"127.0.0.1",
			ServerName:"www.baidu.com",
		},
		Server{
			ServerIP:"127.0.0.2",
			ServerName:"www.qq.com",
		},
	}
	output, err := xml.MarshalIndent(conf, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	//fmt.Println(string(output))
	fmt.Println("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"+string(output))
}


func readXml(fileName string)(conf *Config,err error){
	conf = &Config{}
	data,err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil,err
	}

	err = xml.Unmarshal(data,conf)
	if err != nil {
		return nil,err
	}
	fmt.Println(conf)
	return conf,nil
}

func main(){
	wrieXml()
	/*
	conf,err := readXml("./config.xml")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(conf)
	fmt.Printf("xml: %#v\n",conf)
	*/
}
