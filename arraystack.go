package gopds

// ArrayStack is a persistent stack based on array
type ArrayStack struct {
	top  int
	data []arrayStackEntry
}

// arrayStackEntry internal representation of entry
type arrayStackEntry struct {
	index int
	value interface{}
}

// NewArrayStack returns a pointer to the stack
func NewArrayStack() *ArrayStack {
	s := &ArrayStack{
		data: make([]arrayStackEntry, 1),
	}
	s.data[0] = arrayStackEntry{-1, nil}
	return s
}

// Size returns size of i-th stack
func (s *ArrayStack) Size(i int) int {
	if i > s.top {
		return -1
	}
	size := 0
	for {
		i := s.data[i].index
		if i == -1 {
			return size
		}
		size++
	}
}

// IsEmpty returns true if i-th stack is empty
func (s *ArrayStack) IsEmpty(i int) bool {
	return i > s.top || s.data[i].index == -1
}

// Push places value on the top of i-th stack
// and returns id of new stack
func (s *ArrayStack) Push(i int, value interface{}) int {
	s.top++
	entry := arrayStackEntry{i, value}
	if s.top >= len(s.data) {
		s.data = append(s.data, entry)
	}
	return s.top
}

// Pop removes top value from the i-th stack
func (s *ArrayStack) Pop(i int) (interface{}, bool, int) {
	if s.IsEmpty(i) {
		return nil, false, -1
	}
	k := s.data[i]
	return k.value, true, s.Push(k.index, k.value)
}

// Top returns top value on the i-th stack
func (s *ArrayStack) Top(i int) (interface{}, bool) {
	if s.IsEmpty(i) {
		return nil, false
	}
	return s.data[s.data[i].index].value, true
}
