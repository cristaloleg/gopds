package gopds

// Queue represents persistent queue
type Queue struct {
	head *queueNode
	tail *queueNode
}

// queueNode represents internal node
type queueNode struct {
	value interface{}
	prev  *queueNode
}

// NewQueue returns a pointer to Queue
func NewQueue() *Queue {
	return &Queue{}
}

// Size returns size of a queue
func (q *Queue) Size() int {
	if q.IsEmpty() {
		return 0
	}
	size := 1
	for n := q.tail; n != q.head; n = n.prev {
		size++
	}
	return size
}

// IsEmpty returns true if queue is empty
func (q *Queue) IsEmpty() bool {
	return q.head == nil && q.tail == nil
}

// Push places value on the top of the queue
func (q *Queue) Push(value interface{}) *Queue {
	neww := &queueNode{value, q.tail}
	if q.IsEmpty() {
		return &Queue{neww, neww}
	}
	return &Queue{q.head, neww}
}

// Pop removes top value from the queue
func (q *Queue) Pop() (*Queue, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	n := q.tail
	for ; n != q.head; n = n.prev {
	}
	return &Queue{n, q.tail}, true
}

// Top returns top value in the queue
func (q *Queue) Top() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	return q.head.value, true
}
