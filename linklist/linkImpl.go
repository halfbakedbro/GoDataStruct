package linklist

import (
	"errors"
	"fmt"
)

type node struct {
	value interface{}
	next  *node
}

type LinkList struct {
	Len  int
	head *node
}

func InitList() *LinkList {
	return &LinkList{}
}

func (l *LinkList) Traverse() error {

	if l.head == nil {
		return errors.New("Empty link list")
	}

	currentNode := l.head

	for currentNode != nil {
		fmt.Printf("%v->", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Printf("nil\n")
	return nil
}

func (l *LinkList) Prepend(value interface{}) error {
	n := &node{
		value: value,
	}

	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}
	l.Len++
	return nil
}

func (l *LinkList) Append(value interface{}) error {
	n := &node{
		value: value,
	}

	if l.head == nil {
		l.head = n
	} else {
		curN := l.head

		for curN != nil {
			if curN.next == nil {
				break
			} else {
				curN = curN.next
			}
		}

		curN.next = n
	}
	l.Len++

	return nil
}

func (l *LinkList) PopFront() (interface{}, error) {

	if l.Len <= 0 {
		return nil, errors.New("Empty Link List")
	}

	curHead := l.head
	l.head = curHead.next

	l.Len--
	return curHead, nil
}

func (l *LinkList) PopBack() (interface{}, error) {
	if l.Len <= 0 {
		return nil, errors.New("Empty List")
	}

	curNode := l.head
	var prev *node

	for curNode.next != nil {
		prev = curNode
		curNode = curNode.next
	}

	if prev != nil {
		curNode = prev.next
		prev.next = nil
	} else {
		l.head = nil
	}

	l.Len--
	return curNode, nil
}

func (l *LinkList) Reverse() {
	curr := l.head
	var prev *node
	var next *node

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}
