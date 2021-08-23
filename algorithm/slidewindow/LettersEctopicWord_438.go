package main

import "fmt"

/**
字母异位词, lc 438
*/
func main() {
	str := "cbaebabacd"
	target := "abc"
	anagrams := findAnagrams(str, target)
	fmt.Println("anagrams =", anagrams)
}

func findAnagrams(s string, p string) []int {
	pl := len(p)
	sl := len(s)
	if pl > sl {
		return nil
	}
	var result []int
	m := make(map[byte]int)
	for i := 0; i < pl; i++ {
		m[p[i]]++
	}
	for i1 := 0; i1 < pl; i1++ {
		m[s[i1]]--
	}
	for i := 0; i < sl-pl+1; i++ {
		flag := 0
		if i > 0 {
			m[s[i-1]]++
			m[s[i+pl-1]]--
		}
		for _, v := range m {
			if v != 0 {
				flag = 1
				break
			}
		}
		if flag == 0 {
			result = append(result, i)
		}
	}
	return result
}

func findAnagramsForTwoMap(s string, p string) []int {

	need, window := map[byte]int{}, map[byte]int{}
	for i := range p {
		need[p[i]]++
	}
	left, right := 0, 0
	valid := 0
	indexs := []int{}
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				indexs = append(indexs, left)
			}
			d := s[left]
			left++
			if _, ok := window[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return indexs
}
