package queue

import (
	"testing"
)

func TestEnqueueDequeue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)

	val, err := queue.Dequeue()
	if err != nil || val != 1 {
		t.Errorf("Dequeue() = %d, %v; want 1, nil", val, err)
	}

	val, err = queue.Dequeue()
	if err != nil || val != 2 {
		t.Errorf("Dequeue() = %d, %v; want 2, nil", val, err)
	}

	if _, err := queue.Dequeue(); err == nil {
		t.Errorf("Dequeue() on empty queue; want error, got nil")
	}
}

func TestIsEmpty(t *testing.T) {
	queue := NewQueue()

	if !queue.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}

	queue.Enqueue(1)

	if queue.IsEmpty() {
		t.Errorf("IsEmpty() = true; want false")
	}
}

func TestFront(t *testing.T) {
	queue := NewQueue()

	if _, err := queue.Front(); err == nil {
		t.Errorf("Front() on empty queue; want error, got nil")
	}

	queue.Enqueue(1)
	front, err := queue.Front()
	if err != nil || front != 1 {
		t.Errorf("Front() = %d, %v; want 1, nil", front, err)
	}

	queue.Enqueue(2)
	front, err = queue.Front()
	if err != nil || front != 1 { // Front should still be 1
		t.Errorf("Front() = %d, %v; want 1, nil", front, err)
	}
}
