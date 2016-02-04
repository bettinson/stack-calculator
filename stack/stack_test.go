package stack

import "testing"

func TestStackPushPop(t *testing.T) {
	s := New()
	checkSlice := []int{1, 2, 3, 4, 6}
	for _, i := range checkSlice {
		s.Push(i)
	}
	for i := range checkSlice {
		popped, err := s.Pop()
		if err != nil {
			t.Error(err)
		}
		c := checkSlice[len(checkSlice)-i-1]
		if popped != c {
			t.Errorf("%d does not equal %d", popped, c)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	s := New()
	_, err := s.Pop()
	if err != ErrEmpty {
		t.Errorf("Expected error empty")
	}
}
