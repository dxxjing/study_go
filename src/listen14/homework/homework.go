package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
)
//计算器 原理
// 9+(3-1)*3+10/2
//整成后缀表达是 即：9 3 1 - 3 * 10 2 / + +
//遇到运算符 则从栈中取出最上面两个数 做运算 然后将结果在放入栈中
func main(){
	app := cli.NewApp()
	app.Name = "calc"
	app.Usage = "计算器"
	app.Action = func(c *cli.Context) error {
		return Process(c)
	}
	app.Run(os.Args)
}

func Process(cli *cli.Context) (err error) {
	for {
		//获取终端输入
		reader := bufio.NewReader(os.Stdin)
		express,_ := reader.ReadString('\n')
		express = strings.TrimSpace(express)
		if len(express) == 0 {
			continue
		}

		postExpress,_ := transPost(express)

		res,_ := calc(postExpress)

		fmt.Println(res)
	}
}

func transPost(express string) (postExpress []string,err error){
	var opStack Stack
	var i int
LABEL:
	for i < len(express) {
		switch {
		//35+27
		case express[i] >= '0' && express[i] <= '9':
			var number []byte
			for ; i < len(express); i++ {
				if express[i] < '0' || express[i] > '9' {
					break
				}
				number = append(number, express[i])
			}
			//numStack.Push(string(number))
			postExpress = append(postExpress, string(number))
		case express[i] == '+' || express[i] == '-' || express[i] == '*' || express[i] == '/':
			if opStack.Empty() {
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue LABEL
			}
			data, _ := opStack.Top()
			if data[0] == '(' || data[0] == ')' {
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue LABEL
			}
			if (express[i] == '+' || express[i] == '-') ||
				((express[i] == '*' || express[i] == '/') && (data[0] == '*' || data[0] == '/')) {
				//numStack.Push(data)
				postExpress = append(postExpress, data)
				opStack.Pop()
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue LABEL
			}
			opStack.Push(fmt.Sprintf("%c", express[i]))
			i++
		case express[i] == '(':
			opStack.Push(fmt.Sprintf("%c", express[i]))
			i++
		case express[i] == ')':
			for !opStack.Empty() {
				data, _ := opStack.Pop()
				if data[0] == '(' {
					break
				}
				postExpress = append(postExpress, data)
				//numStack.Push(data)
			}
			i++
		default:
			err = fmt.Errorf("invalid express:%v", express[i])
			return
		}
	}

	for !opStack.Empty() {
		data, _ := opStack.Pop()
		if data[0] == '#' {
			break
		}

		postExpress = append(postExpress, data)
		//numStack.Push(data)
	}
	return
}

func calc(postExpress []string) (result int64,err error) {
	var n1, n2 string
	var s Stack
	for i := 0; i < len(postExpress); i++ {
		var cur = postExpress[i]
		if cur[0] == '-' || cur[0] == '+' || cur[0] == '*' || cur[0] == '/' {
			n1, err = s.Pop()
			if err != nil {
				return
			}
			n2, err = s.Pop()
			if err != nil {
				return
			}

			num2, _ := strconv.Atoi(n1)
			num1, _ := strconv.Atoi(n2)
			var r1 int

			switch cur[0] {
			case '+':
				r1 = num1 + num2
			case '-':
				r1 = num1 - num2
			case '*':
				r1 = num1 * num2
			case '/':
				r1 = num1 / num2
			default:
				err = fmt.Errorf("invalid op")
				return
			}

			s.Push(fmt.Sprintf("%d", r1))
		} else {
			s.Push(cur)
		}
	}

	resultStr, err := s.Top()
	if err != nil {
		return
	}
	result, err = strconv.ParseInt(resultStr, 10, 64)
	return
}