package iniConfig

import (
	"fmt"
	"io/ioutil"
	"testing"
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

func TestIniConfig(t *testing.T){
	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("read file err")
	}
	var config Config
	err = UnMarshal(data,&config)
	if err != nil {
		t.Errorf("unmarshal err:%v\n",err)
	}
	//fmt.Println(config)
	//{{10.238.2.2 8080} {root root test 192.168.1.1 3838 1.2}}
}
