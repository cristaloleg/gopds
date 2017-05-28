package gopds

import "testing"

func TestArrayStack(t *testing.T) {
	s := NewArrayStack()
	if s == nil {
		t.Error("cannot instantiate ArrayStack")
	}

	if !s.IsEmpty(0) || !s.IsEmpty(1) {
		t.Error("want empty")
	}
	if s.Size(0) != 0 || s.Size(1) != -1 {
		t.Error("want empty")
	}
	if value, ok := s.Top(0); ok || value != nil {
		t.Error("want empty")
	}
	if value, ok, next := s.Pop(0); ok || value != nil || next != -1 {
		t.Error("want empty")
	}

	next := s.Push(0, 10)
	if next != 1 {
		t.Errorf("want next %v, got %v", 1, next)
	}

	next = s.Push(1, 20)
	if next != 2 {
		t.Errorf("want next %v, got %v", 2, next)
	}

	next = s.Push(0, 30)
	if next != 3 {
		t.Errorf("want next %v, got %v", 3, next)
	}

	value, ok, next := s.Pop(next)
	if !ok || value != 30 || next != 4 {
		t.Errorf("want %v, got %v, next %v, got %v", 30, value, 4, next)
	}

	if value, ok := s.Top(next); !ok || value != nil {
		t.Errorf("want %v, got %v", nil, value)
	}
}
