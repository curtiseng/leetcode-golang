package main

type Stack struct {
	items []int // 元素
	count int      // 栈中元素个数
	n     int      // 栈的大小
}

func NewStack(n int) *Stack {
	stack := &Stack {
		items: make([]int, n),
		count: 0,
		n:     n,
	}
	return stack
}

func (stack *Stack) Push(item int) bool {
	if  stack.count == stack.n {
		return false
	}
	stack.items[stack.count] = item
	stack.count += 1
	return true
}

func (stack *Stack) Pop() int {
	if stack.count == 0 {
		return -1
	}
	item := stack.items[stack.count - 1]
	stack.count -= 1
	return item
}

func (stack *Stack) Peek() int {
	if stack.count == 0 {
		return -1
	}
	item := stack.items[stack.count - 1]
	return item
}

func (stack *Stack) IsEmpty() bool {
	return stack.count == 0
}