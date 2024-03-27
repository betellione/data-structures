package queue

import (
	"errors"
)

type Queue struct {
	elements []int
}

func NewQueue() *Queue {
	return &Queue{elements: []int{}}
}

func (q *Queue) Enqueue(element int) {
	q.elements = append(q.elements, element)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.elements) == 0 {
		return 0, errors.New("queue is empty")
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Front() (int, error) {
	if len(q.elements) == 0 {
		return 0, errors.New("queue is empty")
	}
	return q.elements[0], nil
}
