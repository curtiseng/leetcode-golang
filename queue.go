package main


type MyCircularDeque struct {
	items []int
	n int
	head int
	tail int
}

func Constructor(n int) MyCircularDeque {
	return MyCircularDeque{
		make([]int, n+1, n+1),  n+1, 0, 0,
	}
}

func (queue *MyCircularDeque) isEmpty() bool {
	if queue.head == queue.tail {
		return true
	}
	return false
}

func (queue *MyCircularDeque) IsEmpty() bool {
	if queue.head == queue.tail {
		return true
	}
	return false
}

func (queue *MyCircularDeque) isFull() bool {
	if (queue.tail + 1) % queue.n == queue.head {
		return true
	}
	return false
}

func (queue *MyCircularDeque) IsFull() bool {
	if (queue.tail + 1) % queue.n == queue.head {
		return true
	}
	return false
}

// 如果head = 0, item放在末尾
// 如果head > 0, item放在前一位置
func (queue *MyCircularDeque) InsertFront(item int) bool {
	if queue.isFull() {
		return false
	}
	queue.head = (queue.n + queue.head - 1) % queue.n
	queue.items[queue.head] = item
	return true
}

func (queue *MyCircularDeque) DeleteFront() bool {
	if queue.isEmpty() {
		return false
	}
	queue.head = (queue.head + 1) % queue.n
	return true
}

func (queue *MyCircularDeque) InsertLast(item int) bool {
	if queue.isFull() {
		return false
	}
	queue.items[queue.tail] = item
	queue.tail = (queue.tail + 1) % queue.n
	return true
}

func (queue *MyCircularDeque) DeleteLast() bool {
	if queue.isEmpty() {
		return false
	}
	queue.tail = (queue.n + queue.tail - 1) % queue.n
	return true
}

func (queue *MyCircularDeque) GetFront() int {
	if queue.isEmpty() {
		return -1
	}
	return queue.items[queue.head]
}

func (queue *MyCircularDeque) GetRear() int {
	if queue.isEmpty() {
		return -1
	}
	return queue.items[(queue.n + queue.tail -1) % queue.n]
}
