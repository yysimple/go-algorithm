package main

import "fmt"

/**
整数转罗马数字,lc 12
*/
func main() {
	roman := integerToRoman(2834)
	fmt.Println("roman: ", roman)
}

func integerToRoman(num int) []string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var res []string
	index := 0
	for index < 13 {
		if num >= nums[index] {
			res = append(res, romans[index])
			num -= nums[index]
		} else {
			index++
		}
	}
	return res
}
