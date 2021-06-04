package queue

import (
	"reflect"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(3)

	eq := reflect.DeepEqual(q.ToSlice(), []interface{}{1, 3})

	if !eq {
		t.Fail()
	}
}

func TestDequeue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Dequeue()

	eq := reflect.DeepEqual(q.ToSlice(), []interface{}{3, 4})

	if !eq {
		t.Fail()
	}
}

func TestSize(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(4)

	eq := q.Size() == 3

	if !eq {
		t.Fail()
	}
}

func TestDequeueOnEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()
	q := Queue{}
	q.Dequeue()
}

func TestIsEmpty(t *testing.T) {
	q := Queue{}
	eq := q.IsEmpty()

	if !eq {
		t.Fail()
	}
}
