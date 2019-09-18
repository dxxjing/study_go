package main

import "fmt"

//定义工资接口
type Cost interface {
	Pay() int
}

//经理
type Manager struct{

}
//技工
type SkillWorker struct{

}
//保洁
type Cleaner struct{

}

func (m Manager) Pay() int {
	return 8600
}

func (s SkillWorker) Pay() int{
	return 5200
}

func (c Cleaner) Pay() int{
	return 3000
}

func main(){
	var (
		eCost Cost
		s SkillWorker
		m Manager
		c Cleaner
	)
	eCost = m
	fmt.Printf("经理：%d\n",eCost.Pay())

	eCost = s
	fmt.Printf("技工：%d\n",eCost.Pay())

	eCost = c
	fmt.Printf("保洁：%d\n",eCost.Pay())
}