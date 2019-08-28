package main

import(
	"fmt"
)

//map 可以嵌套任何类型

//map 切片
func mSlice(){
	var m []map[string]int
	m = make([]map[string]int,10)
	fmt.Printf("map slice %v,%v\n",m,m[0])
	//map slice [map[] map[] map[] map[] map[] map[] map[] map[] map[] map[]],map[]

	for i := 0; i < len(m); i++ {
		if m[i] == nil {
			m[i] = make(map[string]int)
		}
	}
	m[0]["a"] = 1
	m[0]["b"] = 2
	fmt.Println(m)
	//[map[a:1 b:2] map[] map[] map[] map[] map[] map[] map[] map[] map[]]
}

//map: string=>[]int
func mSlice2(){
	m := make(map[string][]int)

	if _,ok := m["a"]; !ok {
		m["a"] = make([]int,10)
	}
	m["a"] = []int{1,2,3}
	fmt.Println(m) //map[a:[1 2 3]]

}

//map嵌套: string=>map[int]int
func mSlice3(keys []string){
	m := make(map[string]map[int]int)

	for _,v := range keys {
		if _,ok := m[v]; !ok {
			m[v] = make(map[int]int)
		}
		m[v][0] = 1
		m[v][1] = 2
	}
	fmt.Println(m)
	//map[a:map[0:1 1:2] b:map[0:1 1:2] c:map[0:1 1:2]]
}

func main(){
	//mSlice()
	//mSlice2()
	mSlice3([]string{"a","b","c"})
}