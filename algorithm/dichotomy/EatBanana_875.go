package main

import (
	"fmt"
	"math"
)

/**
爱吃香蕉的珂珂, lc 875
*/
func main() {
	piles := []int{30, 11, 23, 4, 20}
	hour := 5
	speed := findFitSpeed(piles, hour)
	fmt.Println("speed: ", speed)
}

func findFitSpeed(piles []int, hour int) int {
	if len(piles) == 0 {
		return 0
	}
	// 求出最大值
	maxSpeed := piles[0]
	for _, v := range piles {
		if v > maxSpeed {
			maxSpeed = v
		}
	}
	// 定义最小速度为1，不管怎样，都会去吃香蕉，想到于 left
	minSpeed := 1
	// 条件是 最小速度 大于 最大速度的时候结束，这个时候 minSpeed就是要的最小速度
	for minSpeed <= maxSpeed {
		// 每次都取最新的中间速度
		midSpeed := (maxSpeed + minSpeed) >> 1
		// 这里判断能否在指定时间内吃完，能 就继续去找小的速度，不能 则去找能完成的最小速度
		if canEatComplete(piles, midSpeed, hour) {
			maxSpeed = midSpeed - 1
		} else {
			minSpeed = midSpeed + 1
		}
	}
	return minSpeed
}

/**
这里是求能否在固定时间内吃完，这里的话是向上取整，比如 6个香蕉，吃的速度是 4，则是需要两个小时的，这里要仔细看下题目的意思
*/
func canEatComplete(piles []int, speed int, hour int) bool {
	sum := 0
	for _, v := range piles {
		t := int(math.Ceil(float64(v) / float64(speed)))
		sum += t
	}
	return sum <= hour
}
