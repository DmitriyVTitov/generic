package generic

import "sync"

// Queue is a generic FIFO queue.
type Queue[T any] struct {
	mu       sync.Mutex
	elements []*T
}

// Enqueue adds an object to the end of the queue.
func (q *Queue[T]) Enqueue(el *T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.elements = append(q.elements, el)
}

// Dequeue returns the first object in the queue.
func (q *Queue[T]) Dequeue() *T {
	q.mu.Lock()
	defer q.mu.Unlock()

	var item *T

	if len(q.elements) > 0 {
		item = q.elements[0]
		q.elements[0] = nil // Clear reference to the dequeued item for GC.
		q.elements = q.elements[1:]
	}

	return item
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.elements)
}

// All returns all elements and clears the queue.
func (q *Queue[T]) All() []*T {
	q.mu.Lock()
	defer q.mu.Unlock()

	elements := q.elements
	q.elements = []*T{}

	return elements
}
