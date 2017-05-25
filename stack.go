package gopds

// Stack persistent
type Stack struct {
	top *stackNode
}

// stackNode represents internal node
type stackNode struct {
	value interface{}
	prev  *stackNode
}

// NewStack returns a pointer to Stack
func NewStack() *Stack {
	return &Stack{}
}

// Size returns size of a stack
func (s *Stack) Size() int {
	size := 0
	for n := s.top; n != nil; n = n.prev {
		size++
	}
	return size
}

// IsEmpty returns true if stack is empty
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

// Push places value on the top of stack
func (s *Stack) Push(value interface{}) *Stack {
	n := &stackNode{value, s.top}
	return &Stack{n}
}

// Pop removes top value from the stack
func (s *Stack) Pop() (*Stack, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return &Stack{s.top.prev}, true
}

// Top returns top value on the stack
func (s *Stack) Top() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return s.top.value, true
}
