package linkedList

import (
	"fmt"
)

// Node представляет узел в связном списке
type Node struct {
	Value int
	Next  *Node
}

// LinkedList представляет связный список
type LinkedList struct {
	Head *Node
}

// Append добавляет новый элемент в конец списка
func (ll *LinkedList) Append(value int) {
	newNode := &Node{Value: value}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Prepend добавляет новый элемент в начало списка
func (ll *LinkedList) Prepend(value int) {
	ll.Head = &Node{Value: value, Next: ll.Head}
}

// ToString возвращает строковое представление списка
func (ll *LinkedList) ToString() string {
	var result string
	current := ll.Head
	for current != nil {
		result += fmt.Sprintf("%d -> ", current.Value)
		current = current.Next
	}
	result += "nil"
	return result
}

func main() {
	ll := LinkedList{}
	ll.Append(1)
	ll.Append(2)
	ll.Prepend(0)
	fmt.Println(ll.ToString()) // 0 -> 1 -> 2 -> nil
}
