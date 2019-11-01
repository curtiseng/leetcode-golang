package main

import (
	"fmt"
	"testing"
)


func TestBubbleSort(t *testing.T) {
	a := []int{1, 6, 4, 4, 5}
	n := len(a)
	bubbleSort(a, n)
	//insertionSort(a, n)
	fmt.Printf("a = %v", a)
}