package ds

import (
	"fmt"
	"testing"
)

func TestBinaryTreeInsertSingle(t *testing.T) {
	var tree *BinaryTree
	tree = NewBinaryTree()

	var v interface{}
	var err error
	v, err = tree.Search(1)
	if err == nil {
		t.Errorf("TestBinaryTreeInsertSingle failed: Search() returned nil error for non initialized tree")
	}

	tree.Insert(1, "key1")
	v, err = tree.Search(1)
	if err != nil {
		t.Errorf("TestBinaryTreeInsertSingle failed: Search() returned non nil error")
	}
	if v != "key1" {
		t.Errorf("TestBinaryTreeInsertSingle failed: Search() returned different key")
	}
}

func TestBinaryTreeInsertMultiple(t *testing.T) {
	var tree *BinaryTree
	tree = NewBinaryTree()

	tree.Insert(1, "key1")
	tree.Insert(2, "key2")

	var v interface{}
	var err error
	v, err = tree.Search(1)
	if err != nil {
		t.Errorf("TestBinaryTree_Insert failed: Search() returned non nil error")
	}
	if v != "key1" {
		t.Errorf("TestBinaryTree_Insert failed: Search() returned different key")
	}

	v, err = tree.Search(2)
	if err != nil {
		t.Errorf("TestBinaryTree_Insert failed: Search() returned non nil error")
	}
	if v != "key2" {
		t.Errorf("TestBinaryTree_Insert failed: Search() returned different key")
	}
}

func TestBinaryTreeMaximum(t *testing.T) {
	var tree *BinaryTree
	tree = NewBinaryTree()

	var v interface{}
	var k int
	var err error
	_, _, err = tree.Maximum()
	if err == nil {
		t.Errorf("TestBinaryTreeInsertSingle failed: Maximum() returned nil error for non initialized tree")
	}

	tree.Insert(1, "key1")
	k, v, err = tree.Maximum()
	if err != nil {
		t.Errorf("TestBinaryTreeInsertSingle failed: Maximum() returned non nil error")
	}
	if k != 1 || v != "key1" {
		t.Errorf("TestBinaryTreeInsertSingle failed: Maximum() returned (%v, %v), expected (1, key1)", k, v)
	}

	tree.Insert(2, "key2")

	k, v, err = tree.Maximum()
	if err != nil {
		t.Errorf("TestBinaryTree_Insert failed: Search() returned non nil error")
	}
	if k != 2 || v != "key2" {
		t.Errorf("TestBinaryTreeInsertSingle failed: Maximum() returned (%v, %v), expected (2, key2)", k, v)
	}

	tree.Insert(0, "key0")

	k, v, err = tree.Maximum()
	if err != nil {
		t.Errorf("TestBinaryTree_Insert failed: Search() returned non nil error")
	}
	if k != 2 || v != "key2" {
		t.Errorf("TestBinaryTreeInsertSingle failed: Maximum() returned (%v, %v), expected (2, key2)", k, v)
	}
}

func TestBinaryTreeMinimum(t *testing.T) {
	var tree *BinaryTree
	tree = NewBinaryTree()

	var v interface{}
	var k int
	var err error
	_, _, err = tree.Minimum()
	if err == nil {
		t.Errorf("TestBinaryTreeInsertSingle failed: Minimum() returned nil error for non initialized tree")
	}

	tree.Insert(1, "key1")
	k, v, err = tree.Minimum()
	if err != nil {
		t.Errorf("TestBinaryTreeInsertSingle failed: Minimum() returned non nil error")
	}
	if k != 1 || v != "key1" {
		t.Errorf("TestBinaryTreeInsertSingle failed: Minimum() returned (%v, %v), expected (1, key1)", k, v)
	}

	tree.Insert(2, "key2")

	k, v, err = tree.Minimum()
	if err != nil {
		t.Errorf("TestBinaryTree_Insert failed: Minimum() returned non nil error")
	}
	if k != 1 || v != "key1" {
		t.Errorf("TestBinaryTreeInsertSingle failed: Minimum() returned (%v, %v), expected (1, key1)", k, v)
	}

	tree.Insert(0, "key0")

	k, v, err = tree.Minimum()
	if err != nil {
		t.Errorf("TestBinaryTree_Insert failed: Minimum() returned non nil error")
	}
	if k != 0 || v != "key0" {
		t.Errorf("TestBinaryTreeInsertSingle failed: Minimum() returned (%v, %v), expected (0, key0)", k, v)
	}
}

func TestBinaryTreeSearch(t *testing.T) {
	var tree *BinaryTree
	tree = NewBinaryTree()

	tree.Insert(4, "key4")
	tree.Insert(1, "key1")
	tree.Insert(2, "key2")
	tree.Insert(3, "key3")

	var v interface{}
	var err error

	for i := 1; i <= 4; i++ {
		v, err = tree.Search(i)
		if err != nil {
			t.Errorf("TestBinaryTree_Insert failed: Search() returned non nil error")
		}
		if v != fmt.Sprintf("key%d", i) {
			t.Errorf("TestBinaryTree_Insert failed: Search() returned %v, expected key%d", v, i)
		}
	}
}

func TestBinaryTreeSearchNil(t *testing.T) {
	var tree *BinaryTree
	tree = NewBinaryTree()

	tree.Insert(4, "key4")
	tree.Insert(1, "key1")
	tree.Insert(2, "key2")
	tree.Insert(3, "key3")

	var v interface{}
	var err error

	v, err = tree.Search(0)
	if err != nil || v != nil {
		t.Errorf("TestBinaryTreeSearchNil failed: Search() returned (%v, %v), expected (nil, nil)", v, err)
	}
	v, err = tree.Search(5)
	if err != nil || v != nil {
		t.Errorf("TestBinaryTreeSearchNil failed: Search() returned (%v, %v), expected (nil, nil)", v, err)
	}
}
