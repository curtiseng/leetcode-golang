package main

import "fmt"

type Stack struct {
	items []string // 元素
	count int      // 栈中元素个数
	n     int      // 栈的大小
}

func main() {
	stack := newStack(10)
	stack.push("100")
	stack.push("101")
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	item, err := stack.pop()
	if err != nil {
		fmt.Println("stack is null in main")
	}
	fmt.Println(item)
}

func newStack(n int) *Stack{
	stack := &Stack{
		items: make([]string, n),
		count: 0,
		n:     n,
	}
	return stack
}

func (stack *Stack) push(item string) bool {
	if  stack.count == stack.n {
		return false
	}
	stack.items[stack.count] = item
	stack.count += 1
	return true
}

func (stack *Stack) pop() (item string, err error) {
	if stack.count == 0 {
		err = fmt.Errorf("stack is null")
		return
	}
	item = stack.items[stack.count - 1]
	stack.count -= 1
	return item, err
}