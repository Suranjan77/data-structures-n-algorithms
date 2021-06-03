package linked_list

import (
	"errors"
	"reflect"
)

type linkedList struct {
	head *node
	tail *node
	size int
}

type node struct {
	prev *node
	next *node
	data interface{}
}

func NewLinkedList() *linkedList {
	return &linkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (ll *linkedList) Size() int {
	return ll.size
}

func (ll *linkedList) IsEmpty() bool {
	return ll.size == 0
}

func (ll *linkedList) InsertAtFirst(a interface{}) {
	newNode := &node{prev: nil, next: ll.head, data: a}
	if ll.IsEmpty() {
		ll.tail = newNode
	}
	ll.head = newNode
	ll.size = ll.size + 1
}

func (ll *linkedList) InsertAtLast(a interface{}) {
	newNode := &node{prev: ll.tail, next: nil, data: a}
	ll.tail.next = newNode
	ll.tail = newNode
	ll.size = ll.size + 1
}

func (ll *linkedList) InsertAt(idx int, a interface{}) {
	curr := ll.nodeAt(idx)
	if curr.prev == nil {
		ll.InsertAtFirst(a)
		ll.size = ll.size + 1
	} else if curr.next == nil {
		ll.InsertAtLast(a)
		ll.size = ll.size + 1
	} else {
		newNode := &node{prev: curr.prev, next: curr, data: a}
		curr.prev.next = newNode
		curr.prev = newNode
		ll.size = ll.size + 1
	}
}

func (ll *linkedList) Insert(a interface{}) {
	if ll.head == nil {
		ll.InsertAtFirst(a)
	} else {
		ll.InsertAtLast(a)
	}
}

func (ll *linkedList) RemoveFirst() {
	if !ll.IsEmpty() {
		head := ll.head
		head.next.prev = nil
		ll.head = head.next
		ll.size = ll.size - 1
	}
}

func (ll *linkedList) RemoveLast() {
	if !ll.IsEmpty() {
		tail := ll.head
		for {
			if tail.next == nil {
				break
			}
			tail = tail.next
		}
		tail.prev.next = nil
		ll.tail = tail.prev
		ll.size = ll.size - 1
	}
}

func (ll *linkedList) RemoveAt(idx int) {
	curr := ll.nodeAt(idx)
	if curr.prev == nil {
		ll.RemoveFirst()
	} else if curr.next == nil {
		ll.RemoveLast()
	} else {
		curr.prev.next = curr.next
		curr.next.prev = curr.prev
		ll.size = ll.size - 1
	}
}

func (ll *linkedList) Remove(a interface{}) {
	curr, _ := ll.search(a)
	curr.prev.next = curr.next
	curr.next.prev = curr.prev
	ll.size = ll.size - 1
}

func (ll *linkedList) IndexOf(a interface{}) int {
	_, idx := ll.search(a)
	return idx
}

func (ll *linkedList) PeekFirst() interface{} {
	return ll.head.data
}

func (ll *linkedList) PeekLast() interface{} {
	curr := ll.nodeAt(ll.size)
	return curr.data
}

func (ll *linkedList) PeekAt(idx int) interface{} {
	curr := ll.nodeAt(idx)
	return curr.data
}

func (ll *linkedList) ToSlice() *[]interface{} {
	var sl []interface{}
	head := ll.head
	for {
		if head == nil {
			break
		}
		sl = append(sl, head.data)
		head = head.next
	}
	return &sl
}

func (ll *linkedList) search(a interface{}) (*node, int) {
	head := ll.head
	var found *node
	i := 0
	for {
		if head == nil {
			found = nil
			break
		}
		if reflect.DeepEqual(head.data, a) {
			found = head
			break
		}
		i = i + 1
		head = head.next
	}
	return found, i
}

func (ll *linkedList) nodeAt(idx int) *node {
	err := ll.validateIndex(idx)

	if err != nil {
		panic(err)
	}

	mid := ll.Size() / 2

	if idx == 0 {
		return ll.head
	}
	if idx == ll.Size()-1 {
		return ll.tail
	}

	if idx < mid {
		head := ll.head
		for i := 0; i < mid; i++ {
			if i == idx {
				return head
			}
			head = head.next
		}
		return head
	} else {
		tail := ll.tail
		for i := ll.Size() - 1; i >= mid; i-- {
			if i == idx {
				return tail
			}
			tail = tail.prev
		}
	}

	return nil
}

func (ll *linkedList) validateIndex(idx int) error {
	if idx >= ll.Size() {
		return errors.New("index out of bounds")
	}
	return nil
}
