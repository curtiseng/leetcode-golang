package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	
}

// 交换相邻节点
func swapPairs(head *ListNode) *ListNode {
	preNode := new(ListNode)
	preNode.Next = head
	var resNode = preNode
	for preNode.Next != nil && preNode.Next.Next != nil {
		a, b := preNode.Next, preNode.Next.Next
		preNode.Next, a.Next, b.Next = b, b.Next, a
		preNode = a
	}
	return resNode.Next
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var preNode *ListNode = nil
	currentNode := head
	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = preNode
		preNode = currentNode
		currentNode = nextNode
	}
	return preNode
}