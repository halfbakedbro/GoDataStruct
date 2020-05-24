package queue

type MyQueue struct {
	data []interface{}
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(value interface{}) error {

}
