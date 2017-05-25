package gopds

// Deque represents persistent deque
type Deque struct {
	head  *dequeNode
	child *Deque
	tail  *dequeNode
}

// dequeNode represents internal node
type dequeNode struct {
	value interface{}
	next  *dequeNode
	prev  *dequeNode
}

// NewDeque returns a pointer to Deque
func NewDeque() *Deque {
	return &Deque{}
}

// Size returns size of a stack
func (d *Deque) Size() int {
	if d.IsEmpty() {
		return 0
	}
	size := 0
	if d.head != nil {
		size++
	}
	if d.tail != nil {
		size++
	}
	if d.child != nil {
		size += d.child.Size()
	}
	return size
}

// IsEmpty returns true if stack is empty
func (d *Deque) IsEmpty() bool {
	return d.head == nil && d.tail == nil
}

// PushFront appends element to the front of stack
func (d *Deque) PushFront(value interface{}) *Deque {
	free := &dequeNode{value, d.head, d.tail}
	if d.IsEmpty() {
		return &Deque{free, &Deque{}, nil}
	}
	return &Deque{free, d, nil}
}

// PushBack appends element to the back of stack
func (d *Deque) PushBack(value interface{}) *Deque {
	free := &dequeNode{value, d.head, d.tail}
	if d.IsEmpty() {
		return &Deque{nil, &Deque{}, free}
	}
	return &Deque{nil, d, free}
}

// PopFront removes front value from the stack
func (d *Deque) PopFront() (*Deque, bool) {
	switch {
	case d.IsEmpty():
		return nil, false

	case d.head != nil:
		return &Deque{nil, d.child, d.tail}, true

	default:
		dd, _ := d.child.PopFront()
		return &Deque{nil, dd, d.tail}, true
	}
}

// PopBack removes back value from the stack
func (d *Deque) PopBack() (*Deque, bool) {
	switch {
	case d.IsEmpty():
		return nil, false

	case d.tail != nil:
		return &Deque{d.head, d.child, nil}, true

	default:
		dd, _ := d.child.PopFront()
		return &Deque{d.head, dd, nil}, true
	}
}

// TopFront returns top element from front of the deque
func (d *Deque) TopFront() (interface{}, bool) {
	switch {
	case d.IsEmpty():
		return nil, false
	case d.head != nil:
		return d.head.value, true
	case d.child != nil:
		return d.child.TopFront()
	default:
		return d.tail.value, true
	}
}

// TopBack returns top element from back of the deque
func (d *Deque) TopBack() (interface{}, bool) {
	switch {
	case d.IsEmpty():
		return nil, false
	case d.tail != nil:
		return d.tail.value, true
	case d.child != nil:
		return d.child.TopFront()
	default:
		return d.head.value, true
	}
}
