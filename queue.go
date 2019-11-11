package main

type CircularDeque struct {
	items []int
	n int
	head int
	tail int
}

func NewCircularDeque(n int) CircularDeque {
	return CircularDeque{
		make([]int, n+1, n+1),  n+1, 0, 0,
	}
}

func (queue *CircularDeque) isEmpty() bool {
	if queue.head == queue.tail {
		return true
	}
	return false
}

func (queue *CircularDeque) isFull() bool {
	if (queue.tail + 1) % queue.n == queue.head {
		return true
	}
	return false
}

// 如果head = 0, item放在末尾
// 如果head > 0, item放在前一位置
func (queue *CircularDeque) insertFront(item int) bool {
	if queue.isFull() {
		return false
	}
	queue.head = (queue.n + queue.head - 1) % queue.n
	queue.items[queue.head] = item
	return true
}

func (queue *CircularDeque) deleteFront() bool {
	if queue.isEmpty() {
		return false
	}
	queue.head = (queue.head + 1) % queue.n
	return true
}

func (queue *CircularDeque) insertLast(item int) bool {
	if queue.isFull() {
		return false
	}
	queue.items[queue.tail] = item
	queue.tail = (queue.tail + 1) % queue.n
	return true
}

func (queue *CircularDeque) deleteLast() bool {
	if queue.isEmpty() {
		return false
	}
	queue.tail = (queue.n + queue.tail - 1) % queue.n
	return true
}

func (queue *CircularDeque) getFront() int {
	if queue.isEmpty() {
		return -1
	}
	return queue.items[queue.head]
}

func (queue *CircularDeque) getRear() int {
	if queue.isEmpty() {
		return -1
	}
	return queue.items[(queue.n + queue.tail -1) % queue.n]
}
