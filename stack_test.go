package gopds

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	if s == nil {
		t.Error("cannot instantiate Stack")
	}

	if value, ok := s.Top(); ok || value != nil {
		t.Errorf("expected empty, got %v", value)
	}
	if value := s.Size(); value != 0 {
		t.Errorf("expected empty, got %v", value)
	}

	s1 := s.Push(1)
	if value, ok := s1.Top(); !ok || value != 1 {
		t.Errorf("want %v, got %v", 1, value)
	}
	if value := s1.Size(); value != 1 {
		t.Errorf("expected %v, got %v", 1, value)
	}

	s2, _ := s1.Push(2).Pop()
	if value, ok := s2.Top(); !ok || value != 1 {
		t.Errorf("want %v, got %v", 1, value)
	}
	if value := s2.Size(); value != 1 {
		t.Errorf("expected %v, got %v", 2, value)
	}

	if value, ok := s.Pop(); ok || value != nil {
		t.Error("expected to be empty")
	}
}
