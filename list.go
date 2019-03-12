package main

// 时间复杂度：插入删除O(1)，随机访问O(n)

type ListNode struct {
	Val int
	Next *ListNode // 这里表示Next变量是ListNode型指针
}

func main()  {
	println(5/2)
}

// 链表中环的检测
// 1. 使用Set直接判断是否重复，Set的O(1)时间复杂度，遍历是O(n)
// 2. 快慢指针
func hasCycle(head *ListNode) bool {
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

// 两个有序列表的合并
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	var res *ListNode
	if l1.Val >= l2.Val{
		res = l2
		res.Next = mergeTwoLists(l1,l2.Next)
	}else{
		res = l1
		res.Next = mergeTwoLists(l1.Next,l2)
	}
	return res
}

// 删除链表倒数第n个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preNode := new(ListNode)
	preNode.Next = head
	currentNode := head
	i := 0
	for currentNode != nil {
		i++
		currentNode = currentNode.Next
	}
	i -= n
	currentNode = preNode
	for i > 0 {
		i--
		currentNode = currentNode.Next
	}
	currentNode.Next = currentNode.Next.Next
	return preNode.Next
}

// 求列表的中间结点
// 1.计数
// 2.快慢指针
func middleNode(head *ListNode) *ListNode {
	s := 0
	tempNode := head
	for tempNode != nil {
		s++
		tempNode = tempNode.Next
	}
	for t := 0; t < s / 2; t++ {
		head = head.Next
	}
	return head
}

func middleNode2(head *ListNode) *ListNode {
	slow,fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
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