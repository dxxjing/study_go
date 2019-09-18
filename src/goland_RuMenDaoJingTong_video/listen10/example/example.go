package main

import(
	"fmt"
	"strings"
)

//统计单词数量
func countWord(){
	str := "how do you  do "
	s := strings.Split(str," ")
	fmt.Println(s,len(s))
	m := make(map[string]int,len(s))
	for _,v := range s {
		if v == ""{ //判空不能用" "
			continue
		}
		_,ok := m[v]
		if !ok {
			m[v] = 1
		}else{
			m[v] += 1
		}
	}
	fmt.Printf("countword:%v\n",m)

}

//学生信息的存储 interface{} 可存储任何类型
//student  id name score age
func saveStuInfo(stuMap map[int]map[string]interface{},id,age int,name string,score float32) {
	if _,ok := stuMap[id]; !ok {
		stuMap[id] = make(map[string]interface{})
	}
	stuMap[id]["name"] = name
	stuMap[id]["score"] = score
	stuMap[id]["age"] = age
	//fmt.Println(stuMap)
}


func main(){
	//countWord()


	id1 := 1
	id2 := 2
	//stuMap 类型php的数组
	stuMap := make(map[int]map[string]interface{},10)
	saveStuInfo(stuMap,id1,15,"jdx",98.5)
	saveStuInfo(stuMap,id2,16,"tom",90)
	//可见 map 是引用类型
	fmt.Println(stuMap)
	fmt.Println(stuMap[id1])
	fmt.Println(stuMap[id1]["name"])
	//map[1:map[age:15 name:jdx score:98.5] 2:map[age:16 name:tom score:90]]
	//map[age:15 name:jdx score:98.5]
	//jdx
}
