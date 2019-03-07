package main

// 时间复杂度：插入删除O(1)，随机访问O(n)

type ListNode struct {
	Val int
	Next *ListNode // 这里表示Next变量是ListNode型指针
}

func main()  {

}

// 链表中环的检测

// 两个有序列表的合并

// 删除链表倒数第n个结点

// 求列表的中间结点

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

// 反转单链表，逻辑是依次将后一结点的指针指向前一结点，引入一个值为nil的前置结点作为哨兵，简化逻辑
// 反转过程用到三个指针，分别指向前置、当前、下一个结点
// 先用第三个指针记录下一个结点，然后把当前结点的Next指向前置结点，然后把第一个指针记录到当前结点，第二个指针记录到下一个结点
// 开始下一轮循环
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