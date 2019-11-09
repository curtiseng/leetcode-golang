package main

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	a := []int{1, 1, 2, 2, 3}
	i := removeDuplicates(a)
	//insertionSort(a, n)
	fmt.Printf("i = %v\n", i)
	fmt.Printf("a = %v", a)
}

// 切片是引用类型
func TestRotate(t *testing.T)  {
	nums := []int{1,2,3,4,5,6,7}
	k := 3
	rotate(nums, k)
	fmt.Printf("nums = %v\n", nums)
}