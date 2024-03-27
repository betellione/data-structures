package binaryTree

import "testing"

func TestBinaryTree(t *testing.T) {
	tree := NewBinaryTree(5)

	tree.Root.Insert(3)
	tree.Root.Insert(7)
	tree.Root.Insert(1)

	tests := []struct {
		value    int
		expected bool
	}{
		{3, true},
		{7, true},
		{5, true},
		{1, true},
		{0, false},
		{8, false},
	}

	for _, test := range tests {
		result := tree.Root.Find(test.value)
		if result != test.expected {
			t.Errorf("Find(%d) = %v; want %v", test.value, result, test.expected)
		}
	}
}
