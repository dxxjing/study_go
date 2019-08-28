package main

import (
	"fmt"
	"sort"
)

//仅仅用于排序
//要想排序 必须满足 sort.Interface 才可
//原因：sort.Sort(data interface) 参数data 必须实现sort.Interface 才行
//如果是 int float64 string 切片则无需实现
//sort.Sort(sort.StringSlice(s))即可

//除非是结构体 需要对某个字段排序 否则 无需重写 所以以下是多余的
//对结构体实现 sort.Interface 参见官方文档
/*
type strSlice []string
func (p strSlice) Len() int {
	return len(p)
}
func (p strSlice) Swap(i,j int){
	p[i],p[j] = p[j],p[i]
}
func (p strSlice) Less(i,j int) bool {
	return p[i] > p[j]
}*/

//map 不支持排序 如果要排序 则将指定排序的字段 转为切片 排好序之后 再转为 map
func sortMap(){
	var m  = map[string]int{
		"b":2,
		"a":4,
		"c":6,
	}

	//打算用append 则len 要指定为0 否则会出现 [      a,b,c]
	s := make([]string,0,10)
	//排序
	for k,_ := range m {
		s = append(s,k)
	}
	fmt.Println(s) //[b,a,c]
	//sort.Sort(strSlice(s))
	sort.Sort(sort.StringSlice(s)) //将s 强转为sort.StringSlice 类型
	fmt.Println(s) //[c,b,a]
	newM := make(map[string]int,10)
	for _,v := range s {
		newM[v] = m[v]
	}
	fmt.Println(newM)//map[a:4 b:2 c:6]

}

func main(){
	sortMap()
}
