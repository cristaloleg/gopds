package gopds

import "testing"

func TestDeque(t *testing.T) {
	d := NewDeque()

	if d.Size() != 0 {
		t.Error("expected to be empty")
	}

	d1 := d.PushFront(1)
	if value, ok := d1.TopFront(); !ok || value != 1 {
		t.Errorf("%v %v", 1, value)
	}
	if value := d1.Size(); value != 1 {
		t.Errorf("want size %v, got %v", 1, value)
	}

	d2 := d.PushFront(2)
	if value, ok := d2.TopFront(); !ok || value != 2 {
		t.Errorf("%v %v", 2, value)
	}

	d3 := d1.PushFront(3)
	if value, ok := d3.TopFront(); !ok || value != 3 {
		t.Errorf("%v %v", 3, value)
	}

	d4 := d1.PushBack(4)
	if value, ok := d4.TopBack(); !ok || value != 4 {
		t.Errorf("%v %v", 4, value)
	}

	d5 := d.PushBack(5)
	if value, ok := d5.TopBack(); !ok || value != 5 {
		t.Errorf("%v %v", 5, value)
	}

	d6, ok := d5.PopFront()
	if !ok {
		t.Error("PopFront should be ok")
	}
	if value, ok := d6.TopFront(); !ok || value != 5 {
		t.Errorf("want %v, got %v", 5, value)
	}
}
