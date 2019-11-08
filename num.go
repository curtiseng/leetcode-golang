package main

// 两次遍历，n(n-1)/2，总体来说时间复杂度为O(n^2)
func twoSum(nums []int, target int) []int {
	var out []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				out = append(out, i, j)
			}
		}
	}
	return out
}