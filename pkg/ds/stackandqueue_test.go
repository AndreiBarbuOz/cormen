package ds

import (
	"fmt"
	"testing"
)

func TestSliceStackBasic(t *testing.T) {
	var slice *SliceStack
	slice = NewSliceStack()

	if top := slice.Top(); top != nil {
		t.Errorf("SliceStack init failed: Top() returned non nil")
	}

	val := "test123"
	slice.Push(val)
	if ret := slice.Top(); ret != "test123" {
		t.Errorf("SliceStack init failed: Top() returned non nil")
	}
}

func TestSliceStackEmpty(t *testing.T) {
	slice := NewSliceStack()

	if ret := slice.IsEmpty(); !ret {
		t.Errorf("SliceStack IsEmpty failed: returned false")
	}
	slice.Push("value")
	if ret := slice.IsEmpty(); ret {
		t.Errorf("SliceStack IsEmpty failed: returned true after Push")
	}
	slice.Pop()
	if ret := slice.IsEmpty(); !ret {
		t.Errorf("SliceStack IsEmpty failed: returned true after Push and Pop")
	}
}

func TestSliceStackError(t *testing.T) {
	slice := NewSliceStack()

	slice.Push("value")
	val, err := slice.Pop()
	if val != "value" {
		t.Errorf("SliceStack returned wrong value after Push")
	}
	if err != nil {
		t.Errorf("SliceStack Pop return error on non empty stack")
	}
	val, err = slice.Pop()
	if err == nil {
		t.Errorf("SliceStack Pop return nil error on empty stack")
	}
}

func TestSliceStackAppendPush(t *testing.T) {
	slice := NewSliceStack()

	for i := 0; i < 20; i++ {
		slice.Push(fmt.Sprintf("val%d", i))
	}
	for i := 19; i >= 0; i-- {
		val, err := slice.Pop()
		if val != fmt.Sprintf("val%d", i) {
			t.Errorf("SliceStack returned wrong value after Push")
		}
		if err != nil {
			t.Errorf("SliceStack Pop return error on non empty stack")
		}
	}
}

func TestSliceQueueEmpty(t *testing.T) {
	queue := NewSliceQueue()

	_, err := queue.Dequeue()
	if err == nil {
		t.Errorf("SliceQueue didn't return error when empty")
	}
}

func TestSliceQueueBasic(t *testing.T) {
	queue := NewSliceQueue()

	err := queue.Enqueue("test1")
	if err != nil {
		t.Errorf("SliceQueue didn't Enqueue on empty queue")
	}
	val, err := queue.Dequeue()
	if val != "test1" {
		t.Errorf("SliceQueue didn't Dequeue the correct value: %v\n", val)
	}
	if err != nil {
		t.Errorf("Returned error on Dequeue")
	}
}

func TestSliceQueueOrder(t *testing.T) {
	queue := NewSliceQueue()

	for i := 0; i < 5; i++ {
		err := queue.Enqueue(fmt.Sprintf("val%d", i))
		if err != nil {
			t.Errorf("SliceQueue didn't Enqueue on empty queue")
		}
	}
	for i := 0; i < 5; i++ {
		val, err := queue.Dequeue()
		if err != nil {
			t.Errorf("SliceQueue didn't Enqueue on empty queue")
		}
		if val != fmt.Sprintf("val%d", i) {
			t.Errorf("SliceQueue didn't Dequeue the correct value: %v\n", val)
		}
	}
}

func TestSliceQueueFull(t *testing.T) {
	queue := NewSliceQueue()

	for i := 0; i < 9; i++ {
		err := queue.Enqueue(fmt.Sprintf("val%d", i))
		if err != nil {
			t.Errorf("SliceQueue didn't Enqueue on non-full queue")
		}
	}
	err := queue.Enqueue("val")
	if err == nil {
		t.Errorf("SliceQueue didn't error on full queue")
	}
}

func TestLinkedListBasic(t *testing.T) {
	queue := NewLinkedList()

	err := queue.Enqueue("test1")
	if err != nil {
		t.Errorf("SliceQueue didn't Enqueue on empty queue")
	}
	val, err := queue.Dequeue()
	if val != "test1" {
		t.Errorf("SliceQueue didn't Dequeue the correct value: %v\n", val)
	}
	if err != nil {
		t.Errorf("Returned error on Dequeue")
	}
}

func TestLinkedListOrder(t *testing.T) {
	queue := NewLinkedList()

	for i := 0; i < 5; i++ {
		err := queue.Enqueue(fmt.Sprintf("val%d", i))
		if err != nil {
			t.Errorf("SliceQueue didn't Enqueue on empty queue")
		}
	}
	for i := 0; i < 5; i++ {
		val, err := queue.Dequeue()
		if err != nil {
			t.Errorf("SliceQueue didn't Enqueue on empty queue")
		}
		if val != fmt.Sprintf("val%d", i) {
			t.Errorf("SliceQueue didn't Dequeue the correct value: %v\n", val)
		}
	}
}

func TestLinkedListEmpty(t *testing.T) {
	queue := NewLinkedList()

	_, err := queue.Dequeue()
	if err == nil {
		t.Errorf("LinkedList didn't return error on Dequeue when empty")
	}
	err = queue.Enqueue("test")
	if err != nil {
		t.Errorf("LinkedList didn't Enqueue on empty list")
	}
	_, err = queue.Dequeue()
	if err != nil {
		t.Errorf("LinkedList didn't Dequeue properly")
	}
	_, err = queue.Dequeue()
	if err == nil {
		t.Errorf("LinkedList didn't return error on Dequeue when empty")
	}

}
