package main

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	a := []int{1, 1, 2, 2, 3}
	i := removeDuplicates(a)
	//insertionSort(a, n)
	fmt.Printf("i = %v\n", i)
	fmt.Printf("a = %v", a)
}