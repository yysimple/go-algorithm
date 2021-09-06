package main

import "fmt"

func main() {
	nums := []int{2, 1, 2, 2, 5, 1, 1}
	one_01 := findOnlyOneForMap(nums)
	fmt.Println("res for map：", one_01)
	math := findOnlyOneForMath(nums)
	fmt.Println("res for math：", math)
	bit := findOnlyOneForBit(nums)
	fmt.Println("res for bit：", bit)
	mod := findOnlyOneForMod(nums)
	fmt.Println("res for mod：", mod)

	fmt.Println("---------------------------")

	fmt.Println(int32(5))
	total := int32(0)
	// 0001 0100 -> 0000 1010 & 0000 0001 = 0000 0000
	total += int32(20) >> 1 & 1
	total += int32(20) >> 1 & 1
	// 0000 1010 -> 0000 0101 & 0000 0001 = 0000 0001
	total += int32(10) >> 1 & 1
	total += int32(20) >> 1 & 1
	fmt.Println("total: ", total)
	ans := int32(0)
	ans = ans | 1<<1
	fmt.Println("ans |= 1 << 1: ", ans)
}

/**
map 求解，这种方式就是暴利求解，性能不是很好
*/
func findOnlyOneForMap(nums []int) int {
	m := make(map[int]int)
	for _, k := range nums {
		//如果是其他语言，请注意对应的判空操作！
		m[k]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

/**
利用数学公式计算，既然是三个数字里面找一个只出现过一次的数字，
那么我们去重，然后 * 3的和，再减去之前数组的和，那剩下来的除以2就是要的值
*/
func findOnlyOneForMath(nums []int) int {
	return (sum(set(nums))*3 - sum(nums)) / 2
}

func sum(nums []int) int {
	res := 0
	for _, v := range nums {
		res = res + v
	}
	return res
}

/**
通过map主键唯一的特性过滤重复元素
*/
func set(slc []int) []int {
	var result []int
	// 存放不重复主键
	tempMap := map[int]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		// 加入map后，map长度变化，则元素不重复
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

func findOnlyOneForBit(nums []int) int {
	number, res := 0, 0
	for i := 0; i < 64; i++ {
		//初始化每一位1的个数为0
		number = 0
		for _, k := range nums {
			//通过右移i位的方式，计算每一位1的个数
			number += (k >> i) & 1
		}
		//最终将抵消后剩余的1放到对应的位数上
		res |= (number) % 3 << i
	}
	return res
}

/**
首先全部转成int，那么就是32位；
先解析这个：total += int32(num) >> i & 1
	1. 让每一个数字的二进制都往右移一位，相当于遍历二进制，然后与 0000 0000 0000 0000 0000 0000 0000 0001 进行与操作
	然后这里仔细想想，最后的 total 是不是应该有四种情况：
		- 3 + 3 + 3 +...+ 1
		- 3 + 3 + 3 +...+ 0
		- 0 + 0 + 0 +...+ 1
		- 0 + 0 + 0 +...+ 0
	2. total % 3 这个只可能有两者情况，0 1

再来解析这个：ans |= 1 << i
	1. ans初始化是 0000 0000 0000 0000 0000 0000 0000 0000，上面得知，total % 3 > 0 只有 1 才能进这个逻辑；
	2. 然后 1 << i， | （这个或运算是只要有 1 则为 1）这个相当于在最终结果对应的 i 位置上赋值 1，因为 1 才有意义，最后得到的答案就是最终的值

*/
func findOnlyOneForMod(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
