package stack

import (
	"testing"
)

func TestPushPop(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)

	val, err := stack.Pop()
	if err != nil || val != 2 {
		t.Errorf("Pop() = %d, %v; want 2, nil", val, err)
	}

	val, err = stack.Pop()
	if err != nil || val != 1 {
		t.Errorf("Pop() = %d, %v; want 1, nil", val, err)
	}

	if _, err := stack.Pop(); err == nil {
		t.Errorf("Pop() on empty stack; want error, got nil")
	}
}

func TestIsEmpty(t *testing.T) {
	stack := NewStack()

	if !stack.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}

	stack.Push(1)

	if stack.IsEmpty() {
		t.Errorf("IsEmpty() = true; want false")
	}
}

func TestTop(t *testing.T) {
	stack := NewStack()

	if _, err := stack.Top(); err == nil {
		t.Errorf("Top() on empty stack; want error, got nil")
	}

	stack.Push(1)
	top, err := stack.Top()
	if err != nil || top != 1 {
		t.Errorf("Top() = %d, %v; want 1, nil", top, err)
	}

	stack.Push(2)
	top, err = stack.Top()
	if err != nil || top != 2 {
		t.Errorf("Top() = %d, %v; want 2, nil", top, err)
	}
}
