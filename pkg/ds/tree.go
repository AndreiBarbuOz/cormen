package ds

import "fmt"

type node struct {
	key   int
	data  interface{}
	left  *node
	right *node
}

func (n *node) insert(k int, el interface{}) {
	if n.key >= k {
		if n.left != nil {
			n.left.insert(k, el)
		} else {
			n.left = &node{k, el, nil, nil}
		}
	} else {
		if n.right != nil {
			n.right.insert(k, el)
		} else {
			n.right = &node{k, el, nil, nil}
		}
	}
}

func (n *node) search(k int) interface{} {
	if n.key == k {
		return n.data
	}
	if n.key > k {
		if n.left != nil {
			return n.left.search(k)
		} else {
			return nil
		}
	} else {
		if n.right != nil {
			return n.right.search(k)
		} else {
			return nil
		}
	}
}

type BinaryTree struct {
	root *node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (b *BinaryTree) Insert(k int, el interface{}) {
	if b.root == nil {
		b.root = &node{k, el, nil, nil}
	} else {
		b.root.insert(k, el)
	}
}

func (b *BinaryTree) Search(k int) (interface{}, error) {
	if b.root == nil {
		return nil, fmt.Errorf("called Search() on BinaryTree with nil root")
	}
	return b.root.search(k), nil
}

func (b *BinaryTree) Maximum() (int, interface{}, error) {
	if b.root == nil {
		return 0, nil, fmt.Errorf("called Maximum() on BinaryTree with nil root")
	}
	var ret *node
	ret = b.root
	for ret.right != nil {
		ret = ret.right
	}
	return ret.key, ret.data, nil
}

func (b *BinaryTree) Minimum() (int, interface{}, error) {
	if b.root == nil {
		return 0, nil, fmt.Errorf("called Minimum() on BinaryTree with nil root")
	}
	var ret *node
	ret = b.root
	for ret.left != nil {
		ret = ret.left
	}
	return ret.key, ret.data, nil
}
