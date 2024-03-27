package hashTable

import "hash/fnv"

type node struct {
	key   string
	value string
	next  *node
}

type HashTable struct {
	buckets []*node
	size    int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([]*node, size),
		size:    size,
	}
}

func hash(key string, size int) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	return int(hasher.Sum32()) % size
}

func (h *HashTable) Set(key string, value string) {
	index := hash(key, h.size)

	if h.buckets[index] == nil {
		h.buckets[index] = &node{key: key, value: value}
		return
	}

	currentNode := h.buckets[index]
	for currentNode != nil {
		if currentNode.key == key {
			currentNode.value = value
			return
		}

		if currentNode.next == nil {
			break
		}
		currentNode = currentNode.next
	}

	currentNode.next = &node{key: key, value: value}
}

func (h *HashTable) Get(key string) (string, bool) {
	index := hash(key, h.size)
	currentNode := h.buckets[index]
	for currentNode != nil {
		if currentNode.key == key {
			return currentNode.value, true
		}
		currentNode = currentNode.next
	}

	return "", false
}
