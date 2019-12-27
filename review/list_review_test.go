package review

import "testing"

func TestPrint(t *testing.T) {
	var listNode = New(1)
	listNode.Next = New(3)
	listNode.Next.Next = New(5)
	Print(listNode)
}