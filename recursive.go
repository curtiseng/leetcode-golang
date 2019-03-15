package main

import "fmt"

// 多返回值函数相加问题

// 在空间复杂度上，因为递归调用一次就会在内存栈中保存一次现场数据，考虑这一部分的开销，空间复杂度为O(n)

var (
	depth int
	hasSolved = make(map[int]int)
)

func main()  {
	fmt.Println(f(5))
}

func f(n int) (int) {
	if n == 1 {return 1}
	if n == 2 {return 2}
	depth += 1
	if depth > 1000 {
		return 0
	}
	if hasSolved[n] != 0 {
		return hasSolved[n]
	}
	i := f(n-1) + f(n-2)
	hasSolved[n] = i
	return i
}