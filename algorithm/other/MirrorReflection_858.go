package main

import "fmt"

/**
镜面反射,lc 858
*/
func main() {
	reflection := mirrorReflection(1, 2)
	fmt.Println("min reflection: ", reflection)
}

func mirrorReflection(p, q int) int {
	m, n := p, q
	var r int
	for n > 0 {
		r = m % n
		m = n
		n = r
	}
	if (p/m)%2 == 0 {
		return 2
	} else if (q/m)%2 == 0 {
		return 0
	} else {
		return 1
	}
}
