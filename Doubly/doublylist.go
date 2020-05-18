package doubly

import "errors"

type node struct {
	value interface{}
	next  *node
	prev  *node
}

type Doubly struct {
	len  int
	head *node
	tail *node
}

func initDoubly() *Doubly {
	return &Doubly{}
}

func (d Doubly) Traverse() {

}

func (d *Doubly) Prepend(val interface{}) error {
	n := &node{value: val}

	if d.len <= 0 {
		d.head = n
		d.tail = n
	} else {
		n.next = d.head
		d.head.prev = n
		d.head = n
	}

	d.len++
	return nil
}

func (d *Doubly) Append(val interface{}) error {
	n := &node{value: val}

	if d.len <= 0 {
		d.head = n
		d.tail = n
	} else {
		d.tail.next = n
		n.prev = d.tail
		d.tail = n
	}

	d.len++
	return nil
}

func (d *Doubly) InsertAt(val interface{}, posn uint) error {
	if posn > d.len {
		return errors.New("Out of bound!! intead use prepend/append")
	}

	n := &node{value: val}
	if posn == 0 {
		n.next = d.head
		d.head.prev = n
		d.head = n
	} else {
		curNode := d.head
		curIndex := 0

		for ; curIndex < posn-1; curIndex++ {
			curNode = curNode.next
		}

		n.prev = curNode
		n.next = curNode.next.next
		curNode.next.prev = n
		curNode.next = n

	}
	d.len++
	return nil
}
