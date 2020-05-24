package queue

import (
	"../linklist"
)

type Queue struct {
	list linklist.LinkList
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value interface{}) error {
	return q.list.Prepend(value)
}

func (q *Queue) Dequeue() (interface{}, error) {
	return q.list.PopBack() //This can be optimized by having tail pointer in
	// LinkList data structures
}

func (q Queue) Length() int {
	return q.list.Len
}

func (q Queue) IsEmpty() bool {
	if q.list.Len == 0 {
		return true
	}

	return false
}
