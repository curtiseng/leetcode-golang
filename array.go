package main

func removeDuplicates(nums []int) int {
	var i = 0
	for j:=1; j< len(nums);j++ {
		if  nums[i] == nums [j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i+1
}
