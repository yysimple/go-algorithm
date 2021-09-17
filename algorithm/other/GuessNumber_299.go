package main

import (
	"fmt"
	"strconv"
)

/**
猜数字游戏，lc 299
*/
func main() {
	secret := "1807"
	guess := "7810"
	hint := getHint(secret, guess)
	fmt.Println("hint: ", hint)

	num := guessNum(secret, guess)
	fmt.Println("num: ", num)

}

func getHint(secret string, guess string) string {
	a, b := 0, 0
	mapS, mapG := make([]int, 10), make([]int, 10)
	for i := range secret {
		//注意：这里是获取对应数字的ASCII码
		tmp, charGuess := secret[i], guess[i]
		if tmp == charGuess {
			a++
		} else {
			mapS[tmp-'0']++
			mapG[charGuess-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		//找到重叠的
		b += min(mapS[i], mapG[i])
	}
	//strconv.Itoa : 整数转字符串
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}

func guessNum(secret string, guess string) string {
	bulls, cows := 0, 0
	secretMap, guessMap := make([]int, 10), make([]int, 10)
	for i := range secret {
		secretValue, guessValue := secret[i], guess[i]
		if secretValue == guessValue {
			bulls++
		} else {
			secretMap[secretValue-'0']++
			guessMap[guessValue-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		cows += min(secretMap[i], guessMap[i])
	}
	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
