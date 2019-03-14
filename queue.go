package main

import "fmt"

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

func (queue *CircularQueue) isEmpty() bool {
	if queue.head == queue.tail {
		return true
	}
	return false
}

func (queue *CircularQueue) enqueue(item string) bool {
	if (queue.tail + 1) % queue.n == queue.head {
		return false
	}
	queue.items[queue.tail] = item
	queue.tail = (queue.tail + 1) % queue.n
	return true
}

func (queue *CircularQueue) dequeue() (item string, err error) {
	if queue.isEmpty() {
		return item, fmt.Errorf("queue is empty")
	}
	item = queue.items[queue.head]
	queue.head = (queue.head + 1) % queue.n
	return item, err
}