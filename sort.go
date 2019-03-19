package main


func bubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := range a{
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
			if !flag {
				break
			}
		}
	}
}

func insertionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := range a {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}
