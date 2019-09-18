package listen19

import (
	"testing"
)
/*
	go test 包名， 执行这个包下面的所有测试用例  go test listen19
	go test 测试源文件  执行这个测试源文件中所有的测试用例
	go test -run 执行指定的用例  go test -run TestAdd
 */


//单元测试 go test 加上 -v 会打印日志
func TestAdd(t *testing.T){
	a,b,sum := 3,1,0
	sum = Add(a,b)
	t.Logf("%d + %d = %d",a,b,sum)
	if sum != 4 {
		t.Fatalf("3 + 1 != 4")
	}
}

//基准测试 压力测试  go test -bench flag
func BenchmarkAdd(t *testing.B){
	for i:=0;i < 10; i++ {
		Add(i,i)
	}
}


