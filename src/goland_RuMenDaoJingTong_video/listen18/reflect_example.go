package main

import (
	"fmt"
	"iniConfig"
	"io/ioutil"
)

type Config struct {
	SrvConf ServerConfig	`ini:"server"`
	SqlConf MysqlConfig		`ini:"mysql"`
}

type ServerConfig struct {
	Ip string	`ini:"ip"`
	Port int	`ini:"port"`
}

type MysqlConfig struct {
	UserName string		`ini:"username"`
	Passwd string		`ini:"passwd"`
	DataBase string		`ini:"database"`
	Host string			`ini:"host"`
	Port int			`ini:"port"`
	TimeOut float32		`ini:"timeout"`
}

func main(){
	data, err := ioutil.ReadFile("../iniConfig/config.ini")
	if err != nil {
		fmt.Println("read file err")
	}
	var conf Config
	err = iniConfig.UnMarshal(data,&conf)
	if err != nil {
		fmt.Printf("unmarshal err:%v\n",err)
	}
	fmt.Println("unmarshal:",conf)
}
