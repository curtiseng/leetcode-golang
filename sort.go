package main


// 冒泡排序：两个循环，冒泡交换位置
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


// 插入排序：分为左右排序好和未排序好两个部分，然后遍历排序好的部分，与右边第一个做比较，
// 用数据移动的方式将右边第一个数据放入合适的位置。
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


// 由上到下，先分区，分区后已然总体有序，然后再处理子问题。快速排序
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
