package main

import "math"

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

func rotate(nums []int, k int)  {
	k = k % len(nums)
	for i := 0; i < k; i++  {
		last := nums[len(nums)-1]
		for j := len(nums)-1; j > 0 ; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
}

// 接雨水
// 栈实现和找括号有一定的区别
func trap(height []int) int {
	stack := NewStack(len(height))
	var current, ans = 0, 0
	for current < len(height) {
		for !stack.IsEmpty() && height[current] > stack.Peek() {
			top := stack.Peek()
			stack.Pop()
			if stack.IsEmpty() {
				break
			}
			distance := current - stack.Peek() -1
			min := height[stack.Peek()]
			if height[current] < min {
				min = height[current]
			}
			ans += distance * (min - top)
		}
		stack.Push(current)
		current += 1
	}
}
