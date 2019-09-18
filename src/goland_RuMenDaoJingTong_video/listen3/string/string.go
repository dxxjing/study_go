package main

import (
	"fmt"
)

//字符串逆序 可以对除中文以外的字符串
func reverseStr(str string) (newStr string){
	
	//字符串不能修改 故需要转为切片
	//byte 一个汉字 3 byte
	strSlice := []byte(str)

	strLen := len(str)
	for i := 0; i < strLen / 2; i++ {
		w := strSlice[strLen-i-1]
		t := strSlice[i]
		strSlice[strLen-i-1] = t
		strSlice[i] = w
	}
	newStr = string(strSlice)
	return newStr
}
//字符串逆序 v2  可以对所有字符串逆序（包含中文）
func reverseStrV2(str string) string {
	//rune 一个汉字 1 rune
	strSlice := []rune(str)
	
	strLen := len(strSlice)//此时长度 应取切片的长度
	for i := 0; i < strLen / 2; i++ {
		w := strSlice[strLen-i-1]
		t := strSlice[i]
		strSlice[strLen-i-1] = t
		strSlice[i] = w
	}
	newStr := string(strSlice)
	return newStr

}

func diffByteRune(){
	str := "hello"
	fmt.Printf("str len:%d\n",len(str)) //5 byte

	str2 := "hello,中国"
	fmt.Printf("str2 len:%d\n",len(str2))//12 byte 汉字每个占 3 byte

	byteSlice := []byte(str2)
	fmt.Printf("str2 byte len:%d\n",len(byteSlice))//12 byte

	runeSlice := []rune(str2)
	fmt.Printf("str2 rune len:%d\n",len(runeSlice))//8 rune
	//结论：
	//byte 每个汉字占 3 byte
	//rune 每个汉字占 1 rune
}

func main(){
	var str string
	str = "hello go"
	newStr := reverseStr(str)
	fmt.Printf("v1 after:%s\n",newStr)

	str = "hello 中国"
	newStr = reverseStrV2(str)
	fmt.Printf("v2 after:%s\n",newStr)

	diffByteRune()
}

