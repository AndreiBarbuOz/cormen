package main

import "github.com/AndreiBarbuOz/cormen/pkg/ds"

type Stack interface {
	Top() interface{}
	Push(interface{})
	Pop() (interface{}, error)
	IsEmpty() bool
}

type Queue interface {
	Enqueue(interface{}) error
	Dequeue() (interface{}, error)
}

func main() {
	var list Queue
	list = ds.NewLinkedList()
	list.Enqueue("value")
}
