package linkedlist

import (
	"errors"
)

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	Value        any
	nextNode     *Node
	previousNode *Node
}

func NewList(elements ...interface{}) *List {
	list := &List{}
	for _, el := range elements {
		list.Push(el)
	}
	return list
}

func (n *Node) Next() *Node {
	return n.nextNode
}

func (n *Node) Prev() *Node {
	return n.previousNode
}

func (l *List) Unshift(v interface{}) {
	node := &Node{Value: v}

	if l.head != nil {
		node.nextNode = l.head
		l.head.previousNode = node
		l.head = node
		return
	}

	l.head = node
	l.tail = node
}

func (l *List) Push(v interface{}) {
	node := &Node{Value: v}

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	if l.tail != nil && l.tail != l.head {
		l.tail.nextNode = node
		node.previousNode = l.tail
		l.tail = node
		return
	}

	l.head.nextNode = node
	node.previousNode = l.head
	l.tail = node

}

func (l *List) Shift() (interface{}, error) {

	if l.head == nil {
		return nil, errors.New("empty list")
	}

	if l.head == l.tail {
		value := l.head.Value
		l.head = nil
		l.tail = nil
		return value, nil
	}

	value := l.head.Value
	l.head = l.head.nextNode
	l.head.previousNode = nil

	return value, nil
}

func (l *List) Pop() (interface{}, error) {

	if l.tail == l.head && l.tail != nil {
		value := l.head.Value
		l.tail = nil
		l.head = nil
		return value, nil
	}

	if l.tail != nil {
		node := l.tail
		l.tail = l.tail.previousNode
		l.tail.nextNode = nil
		return node.Value, nil
	}

	return nil, errors.New("empty list")

}

func (l *List) Reverse() {
	if l.head == nil {
		return
	}

	head, tail, node := l.tail, l.head, l.head

	for node != nil {
		node.nextNode, node.previousNode, node = node.previousNode, node.nextNode, node.nextNode
	}

	l.head = head
	l.tail = tail
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
