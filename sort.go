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


// 由上到下排序，先分区，分区后已然总体有序，然后再处理子问题
func quickSort(a []int) {
	separateSort(a, 0, len(a)-1)
}

// 递归，先由上到下分区排序，然后递归，直到总体有序
func separateSort(a []int, start int, end int) {
	if start >= end {
		return
	}
	i := partition(a, start, end)
	separateSort(a, start, i-1)
	separateSort(a, i+1, end)
}

// 分区原地排序，类似插入排序
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



// 由下到上排序，只是分区，然后分区排序，在合并分区时再排序
func MergeSort(a []int) {
	alen := len(a)
	if alen <= 1 {
		return
	}
	mergeSort(a, 0, alen-1)
}

// 递归直到拆分为单个元素，然后回调merge函数，由下到上排序
func mergeSort(a []int, start int, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(a, start, mid)
	mergeSort(a, mid + 1, end)
	merge(a, start, mid, end)
}

// 合并两个有序数组，使用了中间数组，由于数组空间会释放，所以空间复杂度为O(n)
func merge(a []int, start int, mid int, end int) {
	tmpArr := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if a[i] < a[j] {
			tmpArr[k] = a[i]
			i++
		} else {
			tmpArr[k] = a[j]
			j++
		}
	}
	for ; i <= mid; i++ {
		tmpArr[k] = a[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = a[j]
		k++
	}
	copy(a[start:end+1], tmpArr)
}


