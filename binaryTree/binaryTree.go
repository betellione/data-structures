package binaryTree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTree(value int) *BinaryTree {
	return &BinaryTree{Root: &TreeNode{Value: value}}
}

func (t *TreeNode) Insert(value int) {
	if value <= t.Value {
		if t.Left == nil {
			t.Left = &TreeNode{Value: value}
		} else {
			t.Left.Insert(value)
		}
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{Value: value}
		} else {
			t.Right.Insert(value)
		}
	}
}

func (t *TreeNode) Find(value int) bool {
	if t == nil {
		return false
	}
	if value == t.Value {
		return true
	}
	if value < t.Value {
		return t.Left.Find(value)
	}
	return t.Right.Find(value)
}
