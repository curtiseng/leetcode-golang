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

func (queue *CircularDeque) IsEmpty() bool {
	if queue.head == queue.tail {
		return true
	}
	return false
}

func (queue *CircularDeque) IsFull() bool {
	if (queue.tail + 1) % queue.n == queue.head {
		return true
	}
	return false
}

// 如果head = 0, item放在末尾
// 如果head > 0, item放在前一位置
func (queue *CircularDeque) InsertFront(item int) bool {
	if queue.IsFull() {
		return false
	}
	queue.head = (queue.n + queue.head - 1) % queue.n
	queue.items[queue.head] = item
	return true
}

func (queue *CircularDeque) DeleteFront() bool {
	if queue.IsEmpty() {
		return false
	}
	queue.head = (queue.head + 1) % queue.n
	return true
}

func (queue *CircularDeque) InsertLast(item int) bool {
	if queue.IsFull() {
		return false
	}
	queue.items[queue.tail] = item
	queue.tail = (queue.tail + 1) % queue.n
	return true
}

func (queue *CircularDeque) DeleteLast() bool {
	if queue.IsEmpty() {
		return false
	}
	queue.tail = (queue.n + queue.tail - 1) % queue.n
	return true
}

func (queue *CircularDeque) GetFront() int {
	if queue.IsEmpty() {
		return -1
	}
	return queue.items[queue.head]
}

func (queue *CircularDeque) GetRear() int {
	if queue.IsEmpty() {
		return -1
	}
	return queue.items[(queue.n + queue.tail -1) % queue.n]
}
