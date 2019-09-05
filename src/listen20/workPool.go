package main

import (
	"fmt"
	"math/rand"
)

//任务
type Job struct {
	Id int
	Number int
}
//结果
type Result struct {
	job *Job //这里存地址 避免值拷贝
	res int
}

//取一个整数的各位数之和
func calc(job *Job,res chan *Result){
	var sum int
	num := job.Number
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	r := &Result{
		job: job,
		res: sum,
	}
	res <- r
}

func work(jobChan chan *Job,resultChan chan *Result){
	for job := range jobChan {
		calc(job,resultChan)
	}
}

func startWorker(workNum int,jobChan chan *Job,resultChan chan *Result){
	for i := 0; i < workNum; i++ {
		go work(jobChan,resultChan)
	}
}

func print(resultChan chan *Result){
	for res := range resultChan {
		fmt.Printf("%vth---%v---%d\n",res.job.Id,res.job.Number,res.res)
	}
}
//todo 如何执行指定数量的任务 并成功关闭channel 而不出现死锁
//jobChan 中存入 任务job 的地址，避免值拷贝影响性能 结果集channel 亦是如此
func main(){
	//channel 中存储结构体地址
	jobChan := make(chan *Job,5000)
	resultChan := make(chan *Result,5000)
	//done := make(chan bool,1)

	worknum := 128

	startWorker(worknum,jobChan,resultChan)
	//另启一个协程 打印结果
	go print(resultChan)

	var index int
	//不断地产生job
	for {
		num := rand.Int()//产生int型随机数
		index++
		job := &Job{
			Id:     index,
			Number: num,
		}
		jobChan <- job
		//执行指定数量的任务 关闭任务channel
		/*if index == 1000 {
			//done <- true
			close(jobChan)
			break
		}*/
	}

}
