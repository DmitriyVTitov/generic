package generic

import (
	"testing"
)

func TestQ(t *testing.T) {
	q := Q[int]{}

	q.Enqueue(Ptr(1))
	q.Enqueue(Ptr(2))
	q.Enqueue(Ptr(3))

	if q.Len() != 3 {
		t.Errorf("Expected length 3, got %d", q.Len())
	}

	if val := q.Dequeue(); Val(val) != 1 {
		t.Errorf("Expected 1, got %d", Val(val))
	}

	if val := q.Dequeue(); Val(val) != 2 {
		t.Errorf("Expected 2, got %d", Val(val))
	}

	if val := q.Dequeue(); Val(val) != 3 {
		t.Errorf("Expected 3, got %d", Val(val))
	}

	if q.Len() != 0 {
		t.Errorf("Expected length 0 after dequeuing all elements, got %d", q.Len())
	}
}
