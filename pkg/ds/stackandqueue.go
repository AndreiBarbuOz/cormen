package ds

import "fmt"

type SliceStack struct {
	elems []interface{}
	pos   int
}

func (s *SliceStack) Top() interface{} {
	if s.pos >= 0 {
		return s.elems[s.pos]
	} else {
		return nil
	}
}

func (s *SliceStack) Push(el interface{}) {
	s.pos++
	if len(s.elems) > s.pos {
		s.elems[s.pos] = el
	} else {
		s.elems = append(s.elems, el)
	}
}

func (s *SliceStack) Pop() (interface{}, error) {
	if s.pos == -1 {
		return nil, fmt.Errorf("called Pop() on empty SliceStack")
	}
	ret := s.elems[s.pos]
	s.pos--
	return ret, nil
}

func (s *SliceStack) IsEmpty() bool {
	return s.pos == -1
}

func NewSliceStack() *SliceStack {
	return &SliceStack{make([]interface{}, 1), -1}
}

type SliceQueue struct {
	elems    []interface{}
	head     int
	tail     int
	capacity int
}

func (s *SliceQueue) Enqueue(el interface{}) error {
	if (s.tail+1)%len(s.elems) == s.head {
		return fmt.Errorf("could not Enqueue because SliceQueue is full")
	}
	s.elems[s.tail] = el
	s.tail = (s.tail + 1) % len(s.elems)
	return nil
}

func (s *SliceQueue) Dequeue() (interface{}, error) {
	if s.head == s.tail {
		return nil, fmt.Errorf("called Dequeue on empty SliceQueue")
	}
	ret := s.elems[s.head]
	s.head = (s.head + 1) % len(s.elems)
	return ret, nil
}

func NewSliceQueue() *SliceQueue {
	return &SliceQueue{make([]interface{}, 10), 0, 0, 10}
}

type linkedListElem struct {
	data interface{}
	next *linkedListElem
}

type LinkedList struct {
	head *linkedListElem
	tail *linkedListElem
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Enqueue(el interface{}) error {
	if l.head != nil {
		l.tail.next = &linkedListElem{el, nil}
		l.tail = l.tail.next
	} else {
		l.head = &linkedListElem{el, nil}
		l.tail = l.head
	}
	return nil
}

func (l *LinkedList) Dequeue() (interface{}, error) {
	if l.head == nil {
		return nil, fmt.Errorf("called Dequeue on empty list")
	}
	ret := l.head.data
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	return ret, nil
}

type LinkedStack struct {
	top *linkedListElem
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{}
}

func (l *LinkedStack) Top() interface{} {
	if l.top != nil {
		return l.top.data
	} else {
		return nil
	}
}

func (l *LinkedStack) Push(el interface{}) {
	if l.top != nil {
		newElem := &linkedListElem{el, l.top}
		l.top = newElem
	} else {
		l.top = &linkedListElem{el, nil}
	}
}

func (l *LinkedStack) IsEmpty() bool {
	return l.top == nil
}

func (l *LinkedStack) Pop() (interface{}, error) {
	if l.top == nil {
		return nil, fmt.Errorf("called Pop() on empty LinkedStack")
	}
	var ret interface{}
	ret = l.top.data
	l.top = l.top.next
	return ret, nil
}
