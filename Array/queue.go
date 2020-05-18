package queue

import "errors"

type Queue struct {
	length int
	data   []interface{}
}

func (q *Queue) Enqueue(value interface{}) {
	q.data = append(q.data, value)
	q.length++
}

func (q *Queue) Dequeue() (value interface{}, err error) {
	if q.length <= 0 {
		return nil, errors.New("Queue is empty")
	}

	value = q.data[0]
	q.data = q.data[1:]
	q.length--
	return value, nil
}

func (q Queue) length() int {
	return q.length
}
