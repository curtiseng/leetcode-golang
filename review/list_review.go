package review

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func New(val int) *ListNode {
	return &ListNode{
		Val: val,
		Next: nil,
	}
}

func Print(head *ListNode) {
	fmt.Print(head.Val)
	for head.Next != nil {
		fmt.Print("->", head.Next.Val)
		head = head.Next
	}
}

func hasCycle(head *ListNode) bool{
	if head == nil || head.Next == nil {
		return false
	}
	set := make(map[*ListNode]bool)
	for head != nil {
		if set[head] == true {
			return true
		}
		set[head] = true
		head = head.Next
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
