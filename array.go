package main

/**
 * 前指针记录已完成的不重复item，后指针记录下一个不重复值
 * go语言切片是引用传递，但使用append等函数后可能重新分配地址
 */
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var i = 0
	for j:=1; j< len(nums);j++ {
		if  nums[i] != nums [j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i+1
}
