package gopds

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Error("cannot instantiate Queue")
	}

	if !q.IsEmpty() || q.Size() != 0 {
		t.Error("expected to be empty")
	}
	if value, ok := q.Pop(); ok || value != nil {
		t.Error("expected to be empty")
	}
	if value, ok := q.Top(); ok || value != nil {
		t.Error("expected to be empty")
	}

	q1 := q.Push(1)
	if value, ok := q1.Top(); !ok || value != 1 {
		t.Errorf("want %v, got %v", 1, value)
	}
	if q1.IsEmpty() || q1.Size() != 1 {
		t.Error("expected to be non-empty")
	}

	q2 := q1.Push(2)
	if value, ok := q2.Top(); !ok || value != 1 {
		t.Errorf("want %v, got %v", 1, value)
	}
	if q2.IsEmpty() || q2.Size() != 2 {
		t.Error("expected to be non-empty")
	}

	q3, ok := q2.Pop()
	if value := q3.Size(); !ok || value != 2 {
		t.Errorf("want size %v, got %v", 2, value)
	}
	if value, ok := q3.Top(); !ok || value != 1 {
		t.Errorf("want %v, got %v", 1, value)
	}
}
