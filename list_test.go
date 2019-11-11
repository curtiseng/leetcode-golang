package main

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	var l1 = New(1)
	l1.Next = New(3)
	l1.Next.Next = New(5)
	var l2 = New(2)
	l2.Next = New(4)
	l2.Next.Next = New(6)

	res := mergeTwoLists(l1, l2)

	for {
		fmt.Print(res.Val)
		if res.Next == nil {
			break
		}
		res = res.Next
	}
}
