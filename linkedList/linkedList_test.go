package linkedList

import (
	"testing"
)

func TestAppend(t *testing.T) {
	ll := LinkedList{}
	ll.Append(1)
	if ll.Head.Value != 1 {
		t.Errorf("Expected head value 1, got %d", ll.Head.Value)
	}

	ll.Append(2)
	if ll.Head.Next.Value != 2 {
		t.Errorf("Expected second value 2, got %d", ll.Head.Next.Value)
	}
}

func TestPrepend(t *testing.T) {
	ll := LinkedList{}
	ll.Prepend(1)
	if ll.Head.Value != 1 {
		t.Errorf("Expected head value 1, got %d", ll.Head.Value)
	}

	ll.Prepend(0)
	if ll.Head.Value != 0 {
		t.Errorf("Expected new head value 0, got %d", ll.Head.Value)
	}
}

func TestToString(t *testing.T) {
	ll := LinkedList{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	expected := "1 -> 2 -> 3 -> nil"
	if ll.ToString() != expected {
		t.Errorf("Expected string '%s', got '%s'", expected, ll.ToString())
	}
}
