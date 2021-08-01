package main

import (
	"fmt"
)

/**
买卖股票最佳时机，lc 122
题目：如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。注意你不能在买入股票前卖出股票。
分析：
其实就是要获取，两数相减的是正数的值；1-5，3-6，4-8；然后累加即可；

*/
func main() {

	stockPrices := []int{7, 1, 5, 3, 6, 4, 8, 5, 9}
	profit := maxProfit(stockPrices)
	fmt.Println("profit =", profit)

	add := findAdd(stockPrices)
	fmt.Println("add =", add)

}

func maxProfit(stockPrices []int) int {
	// 1. 判断股票是否支持买卖
	if len(stockPrices) < 2 {
		return 0
	}
	// 2. 初始化数组，和第一天数据 第一个索引代表的是 天数， 第二个索引代表的是 买卖操作，0=卖出，1=买入
	dayAndSellOrBuy := make([][2]int, len(stockPrices))
	// 第1天卖出
	dayAndSellOrBuy[0][0] = 0
	// 第一天买入
	dayAndSellOrBuy[0][1] = -stockPrices[0]
	// 3. 求出每天买卖剩余的收益
	for i := 1; i < len(stockPrices); i++ {
		/**
		1. 卖出之后剩余的钱 = max(前一天卖出剩余的钱, 前一天买入之后剩余的钱 + 当天的卖出股票可以赚钱的钱)
		2. 当天卖出的条件：前一天必须买入；
		3. dayAndSellOrBuy[i-1][1]+stockPrices[i] 这个就是条件，代表前一天是买入操作；
		+stockPrices[i]，就是今天卖出赚到的钱和昨天买入剩余的钱相加，就是当前卖出操作获取的收益
		4. 然后就再 前一天卖出、和当天卖出；这两种情况下收益最高的
		5. 这就是在高价，也就是马上要跌价之前卖出
		*/
		dayAndSellOrBuy[i][0] = max(dayAndSellOrBuy[i-1][0], dayAndSellOrBuy[i-1][1]+stockPrices[i])

		/**
		1. 买入之后剩余的钱 = max(前一天卖出剩余的钱 - 当天买入所用的钱，前一天买入剩下的钱)
		2. 当天买入的条件：前一天必须卖出
		3. dayAndSellOrBuy[i-1][0]-stockPrices[i] 这个代表前一天是卖出操作；
		-stockPrices[i]，表示昨天卖出剩余的钱-今天买入需要花的钱，就是当前买入后持有的收益
		4. 这是就是当天买入，和前一天买入；去最后持有收益最高的情况；
		5. 这里就是在前后几天中的最低价买入；
		*/
		dayAndSellOrBuy[i][1] = max(dayAndSellOrBuy[i-1][0]-stockPrices[i], dayAndSellOrBuy[i-1][1])
	}
	return dayAndSellOrBuy[len(stockPrices)-1][0]
}

/**
这种方式就是最简单的求增加的区域总和，不算股票的买卖场景
*/
func findAdd(arr []int) int {
	profit := 0
	for i := 0; i < len(arr)-1; i++ {
		num := arr[i+1] - arr[i]
		if num > 0 {
			profit += num
		}
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
