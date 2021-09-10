package main

import (
	"fmt"
	"sort"
)

/**
取暖器，lc 475
*/
func main() {
	houses := []int{1, 2, 3, 4, 5, 6}
	heaters := []int{2, 5}
	radius := findRadius(houses, heaters)
	fmt.Println("radius: ", radius)
}

/**
查找最小半径
*/
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	res := 0
	for _, house := range houses {
		left := 0
		right := len(heaters) - 1
		houseRes := 0
		// 二分查找
		for left < right {
			middle := left + (right-left)/2
			if heaters[middle] < house {
				left = middle + 1
			} else {
				right = middle
			}
		}

		if heaters[left] == house {
			houseRes = 0
		} else if heaters[left] < house { // 供暖器在左侧
			houseRes = house - heaters[left]
		} else { // 供暖器在右侧
			if left == 0 {
				houseRes = heaters[left] - house
			} else {
				houseRes = getMin(heaters[left]-house, house-heaters[left-1])
			}
		}

		res = getMax(res, houseRes)
	}

	return res
}

func getMin(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
