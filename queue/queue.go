package queue

import "sync"

type Queue struct {
	items *[]interface{}
	m     sync.Mutex
}

func (q *Queue) IsEmpty() bool {
	if q.items == nil {
		return true
	}
	return len(*q.items) == 0
}

func (q *Queue) Size() int {
	return len(*q.items)
}

func (q *Queue) Enqueue(a interface{}) {
	q.m.Lock()
	defer q.m.Unlock()
	if q.items == nil {
		q.items = &[]interface{}{a}
	} else {
		*q.items = append(*q.items, a)
	}
}

func (q *Queue) Dequeue() {
	if q.IsEmpty() {
		panic("dequeue on empty queue not allowed")
	}
	q.m.Lock()
	defer q.m.Unlock()

	*(*q).items = (*q.items)[1:]
}

func (q *Queue) ToSlice() []interface{} {
	q.m.Lock()
	defer q.m.Unlock()
	return *q.items
}
