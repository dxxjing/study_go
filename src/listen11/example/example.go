package main

import(
	"fmt"
)

/*
	练习：
	名字中包含a或A: 1枚金币
	名字中包含e或E: 1枚金币
	名字中包含i或I: 2枚金币
	名字中包含o或O: 3枚金币
	名字中包含u或U: 5枚金币
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
)

func saveCoin(distribution map[string]int) int {
	useCoins := 0
	for _,name := range users {
		if _,ok := distribution[name];!ok {
			distribution[name] = 0
		}
		for _,ch := range name {
			switch ch{
			case 'a','A','e','E':
				distribution[name] += 1
				useCoins += 1
			case 'i','I':
				distribution[name] += 2
				useCoins += 2
			case 'o','O':
				distribution[name] += 3
				useCoins += 3
			case 'u','U':
				distribution[name] += 5
				useCoins += 5
			}
		}
	}
	c := coins - useCoins
	return c
}

func main(){
	distribution := make(map[string]int)//存储金币
	c := saveCoin(distribution)
	fmt.Println(distribution)
	fmt.Println(c)
	//map[Aaron:5 Adriano:7 Augustus:16 Elizabeth:5 Emilie:6 Giana:4 Heidi:5 Matthew:2 Peter:2 Sarah:2]
	//-4
}
