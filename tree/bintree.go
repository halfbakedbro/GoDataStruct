package tree

import (
	"errors"
	"strings"
)

type node struct {
	val   interface{}
	left  *node
	right *node
}

type BinTree struct {
	root *node
	len  int
}

func checkGreater(val, val1 interface{}) (bool, error) {

	switch v := val.(type) {
	case int:
		if v > val1.(int) {
			return true, nil
		} else if v < val1.(int) {
			return false, nil
		}
	case string:
		x := strings.Compare(v, val1.(string))
		if x == -1 {
			return false, nil
		} else if x == 1 {
			return true, nil
		}
	default:
		return false, errors.New("unknown type")
	}

	return false, nil
}

func initBinTree() *BinTree {
	return &BinTree{root: &node{}}
}

func (n *node) Insert(val interface{}) error {

	if n == nil {
		return errors.New("Empty node wrong pass")
	}

	if ok, err := checkGreater(n.val, val); ok {
		if n.left == nil {
			n.left = &node{val: val}
			return err
		} else {
			n.left.Insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &node{val: val}
			return err
		} else {
			n.right.Insert(val)
		}
	}

	return nil
}

func (b *BinTree) Insert(val interface{}) error {

	if b == nil {
		node := &node{val: val}
		b.root = node
	} else { //error handling required
		b.root.Insert(val)
	}

	return nil
}

func (n *node) Find(val interface{}) (interface{}, bool) {

	if n == nil {
		return 0, false
	} else if n.val == val {
		return val, true
	} else if ok, _ := checkGreater(n.val, val); ok {
		n.right.Find(val)
	} else {
		n.left.Find(val)
	}

	return 0, false
}

func (b BinTree) Find(val interface{}) (interface{}, bool) {
	if b.root == nil {
		return 0, false
	} else {
		return b.root.Find(val)
	}
}

func (b *BinTree) Delete(val interface{}) bool {

	b.root.Delete(val)
	return true
}

func MinVal(n *node) *node {
	curr := n

	for curr != nil && curr.left != nil {
		curr = curr.left
	}

	return curr
}

func (n *node) Delete(val interface{}) *node {
	if n == nil {
		return n
	}

	if ok, _ := checkGreater(n.val, val); ok {
		n.right = n.right.Delete(val)
	} else if !ok {
		n.left = n.left.Delete(val)
	} else /* Same key this is the node to be deleted  */ {
		if n.left == nil { //node with only one child
			temp := n.right
			n = nil
			return temp
		} else if n.right == nil {
			temp := n.left
			n = nil
			return temp
		}

		//node with two children

		temp := MinVal(n.right)

		n.val = temp.val

		n.right = n.right.Delete(temp.val)
	}

	return n
}
