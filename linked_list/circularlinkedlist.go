package linked_list

import (
	"errors"
	"reflect"
)

type circularlinkedList struct {
	head *circularNode
	size int
}

type circularNode struct {
	prev *circularNode
	next *circularNode
	data interface{}
}

func NewCircularLinkedList() *circularlinkedList {
	return &circularlinkedList{
		head: nil,
		size: 0,
	}
}

func (ll *circularlinkedList) Size() int {
	return ll.size
}

func (ll *circularlinkedList) IsEmpty() bool {
	return ll.size == 0
}

func (ll *circularlinkedList) InsertAtFirst(a interface{}) {
	newNode := &circularNode{data: a}
	if ll.IsEmpty() {
		newNode.next = newNode
		newNode.prev = newNode
		ll.head = newNode
	} else {
		newNode.prev = ll.head.prev
		newNode.next = ll.head
		ll.head.prev.next = newNode
		ll.head.prev = newNode
		ll.head = newNode
	}
	ll.size = ll.size + 1
}

func (ll *circularlinkedList) InsertAtLast(a interface{}) {
	newNode := &circularNode{prev: ll.head.prev, next: ll.head, data: a}
	ll.head.prev.next = newNode
	ll.head.prev = newNode
	ll.size = ll.size + 1
}

func (ll *circularlinkedList) InsertAt(idx int, a interface{}) {
	ll.validateIndex(idx)
	curr := ll.nodeAt(idx)
	if curr == ll.head {
		ll.InsertAtFirst(a)
	} else if curr == ll.head.prev {
		ll.InsertAtLast(a)
	} else {
		newNode := &circularNode{prev: curr.prev, next: curr, data: a}
		curr.prev.next = newNode
		curr.prev = newNode
		ll.size = ll.size + 1
	}
}

func (ll *circularlinkedList) Insert(a interface{}) {
	if ll.head == nil {
		ll.InsertAtFirst(a)
	} else {
		ll.InsertAtLast(a)
	}
}

func (ll *circularlinkedList) RemoveFirst() {
	if !ll.IsEmpty() {
		head := ll.head
		head.next.prev = head.prev
		head.prev.next = head.next
		ll.head = head.next
		ll.size = ll.size - 1
	}
}

func (ll *circularlinkedList) RemoveLast() {
	if !ll.IsEmpty() {
		tail := ll.head.prev
		ll.head.prev = tail.prev
		tail.prev.next = ll.head
		ll.size = ll.size - 1
	}
}

func (ll *circularlinkedList) RemoveAt(idx int) {
	if ll.head == nil {
		panic("Remove from empty list not allowed")
	}
	curr := ll.nodeAt(idx)
	if curr == ll.head {
		ll.RemoveFirst()
	} else if curr == ll.head.prev {
		ll.RemoveLast()
	} else {
		curr.prev.next = curr.next
		curr.next.prev = curr.prev
		ll.size = ll.size - 1
	}
}

func (ll *circularlinkedList) Remove(a interface{}) {
	if ll.head == nil {
		panic("Remove from empty list not allowed")
	}
	curr, idx := ll.search(a)
	if curr != nil {
		ll.RemoveAt(idx)
	}
}

func (ll *circularlinkedList) IndexOf(a interface{}) int {
	_, idx := ll.search(a)
	return idx
}

func (ll *circularlinkedList) PeekFirst() interface{} {
	return ll.head.data
}

func (ll *circularlinkedList) PeekLast() interface{} {
	curr := ll.head.prev
	return curr.data
}

func (ll *circularlinkedList) PeekAt(idx int) interface{} {
	curr := ll.nodeAt(idx)
	return curr.data
}

func (ll *circularlinkedList) ToSlice() *[]interface{} {
	var sl []interface{}
	start := ll.head
	next := start
	for {
		sl = append(sl, next.data)
		next = next.next
		if next == start {
			break
		}
	}
	return &sl
}

func (ll *circularlinkedList) search(a interface{}) (*circularNode, int) {
	start := ll.head
	if start == nil {
		return nil, -1
	}
	next := start
	var found *circularNode
	i := 0
	for {
		if reflect.DeepEqual(next.data, a) {
			found = next
			break
		}
		i = i + 1
		next = next.next
		if next == start {
			break
		}
	}
	return found, i
}

func (ll *circularlinkedList) nodeAt(idx int) *circularNode {
	err := ll.validateIndex(idx)

	if err != nil {
		panic(err)
	}

	mid := ll.Size() / 2

	if idx == 0 {
		return ll.head
	}
	if idx == ll.Size()-1 {
		return ll.head.prev
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
		tail := ll.head.prev
		for i := ll.Size() - 1; i >= mid; i-- {
			if i == idx {
				return tail
			}
			tail = tail.prev
		}
	}

	return nil
}

func (ll *circularlinkedList) validateIndex(idx int) error {
	if idx >= ll.Size() {
		return errors.New("index out of bounds")
	}
	return nil
}
