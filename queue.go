package main

type CircularQueue struct {
	items []string
	n int
	head int
	tail int
}

func newCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue {
		make([]string, n), n, 0, 0,
	}
}

func (queue *CircularQueue) IsEmpty() bool {
	if queue.head == queue.tail {
		return true
	}
	return false
}