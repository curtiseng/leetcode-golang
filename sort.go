package main



func bubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := range a{
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
			if !flag {
				break
			}
		}
	}
}



func insertionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := range a {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}


// 由上到下，快速排序
func quickSort(a []int) {
	separateSort(a, 0, len(a)-1)
}
func separateSort(a []int, start int, end int) {
	if start >= end {
		return
	}
	i := partition(a, start, end)
	separateSort(a, start, i-1)
	separateSort(a, i+1, end)
}
func partition(a []int, start int, end int) int {
	// 选取最后一位当对比数字
	pivot := a[end]
	var i = start
	for j := start; j < end; j++ {
		if a[j] < pivot {
			if !(i == j) {
				// 交换位置
				a[i], a[j] = a[j], a[i]
			}
			i++
		}
	}
	a[i], a[end] = a[end], a[i]
	return i
}
