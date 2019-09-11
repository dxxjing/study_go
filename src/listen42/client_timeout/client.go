package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Respack struct{
	r *http.Response
	err error
}

var wg *sync.WaitGroup

func work(ctx context.Context){
	defer wg.Done()

	tr := &http.Transport{}
	cli := &http.Client{
		Transport:tr,
	}
	c := make(chan Respack,1)
	req,_ := http.NewRequest("GET","http://localhost:9090",nil)
	//新开goroutine 请求http server
	go func(){
		resp,err := cli.Do(req)
		pack := Respack{r:resp,err:err}
		c <- pack
	}()

	select{
	case <-ctx.Done():
		//tr.CancelRequest() 该函数已废弃
		<-c
		fmt.Println("time out")
		return
	case res := <-c:
		if res.err != nil {
			fmt.Println(res.err)
			return
		}
		defer res.r.Body.Close()
		out,_ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("server response:%s",string(out))
	}
}

func main(){

	ctx,cancel := context.WithTimeout(context.Background(),2 * time.Second)
	defer cancel()

	wg.Add(1)
	go work(ctx)
	wg.Wait()

	fmt.Println("finished")
}
