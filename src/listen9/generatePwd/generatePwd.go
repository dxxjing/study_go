package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

/*
	命令行模式 生成密码
	-l 20 指定密码长度
	-t 
		num 全数字密码
		char 全英文
		mix 包含数字和英文,特殊符号
*/

var (
	pwdLength int
	pwdType string
)

const (
	numStr = "0123456789"
 	charStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
 	specStr = "&*@#$%()[]{}?"
)

func parseParm(){
	//-l -t 绑定时不要加-,flag 会自动帮你加 否则就变成 --t --l
	flag.IntVar(&pwdLength,"l",32,"指定密码长度")
	flag.StringVar(&pwdType,"t","mix",`可选参数:num 全数字 char 全英文 mix 包含数字和英文,特殊符号`)
	flag.Parse()
}

func genPassword() string {
	var (
		newStr string
		pwdSlice = make([]byte,pwdLength,pwdLength)
	)
	if pwdType == "num" {
		newStr = numStr
	}else if pwdType == "char" {
		newStr = charStr
	}else{
		newStr = numStr + charStr + specStr
	}
	newStrLen := len(newStr)
	for i := 0; i < pwdLength; i++ {
		index := rand.Intn(newStrLen) //返回一个取值范围在[0,newStrLen)的伪随机int值
		pwdSlice[i] = newStr[index]
	}
	return string(pwdSlice)
}

func main(){

	parseParm()
	//命令：generatePwd.exe -l 64 -t char
	fmt.Printf("pwdLength:%d,pwdType:%s \n",pwdLength,pwdType)

	//随机数种子 否则会产生重复的随机数
	rand.Seed(time.Now().UnixNano())
	pwd := genPassword()
	fmt.Printf("gen pwd %s \n",pwd)
}
