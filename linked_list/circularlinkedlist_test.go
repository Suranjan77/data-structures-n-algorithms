package linked_list

import (
	"reflect"
	"testing"
)

func TestCircInsertAtFirst(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.InsertAtFirst(6)
	eq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{6})
	linkedList.InsertAtFirst(7)
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{7, 6})
	siz := linkedList.Size() == 2
	if !eq || !seq || !siz {
		t.Fail()
	}
}

func TestCircInsertAtZero(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.InsertAt(0, 5)
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{5, 6, 7})
	siz := linkedList.Size() == 3
	if !seq || !siz {
		t.Fail()
	}
}

func TestCircInsertAtEnd(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.InsertAt(1, 5)
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{6, 7, 5})
	siz := linkedList.Size() == 3
	if !seq || !siz {
		t.Fail()
	}
}

func TestCircRemoveEmptyList(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	linkedList := NewCircularLinkedList()
	linkedList.Remove(5)
}

func TestCircRemoveOutOfIndex(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	linkedList := NewCircularLinkedList()
	linkedList.Insert(10)
	linkedList.Insert(50)
	linkedList.RemoveAt(10)
}

func TestCircInsert(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(8)
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{6, 8})
	siz := linkedList.Size() == 2
	if !seq || !siz {
		t.Fail()
	}
}

func TestCircInsertAtLast(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.InsertAtLast(5)
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{6, 7, 5})
	siz := linkedList.Size() == 3
	if !seq || !siz {
		t.Fail()
	}
}

func TestCircInsertAt(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.Insert(5)
	linkedList.InsertAt(1, 10)
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{6, 10, 7, 5})
	siz := linkedList.Size() == 4
	if !seq || !siz {
		t.Fail()
	}
}

func TestCircRemoveFirst(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.Insert(5)
	linkedList.Insert(10)
	linkedList.RemoveFirst()
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{7, 5, 10})
	siz := linkedList.Size() == 3
	if !seq || !siz {
		t.Fail()
	}
}

func TestCircRemoveLast(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.Insert(5)
	linkedList.Insert(10)
	linkedList.RemoveLast()
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{6, 7, 5})
	siz := linkedList.Size() == 3
	if !seq || !siz {
		t.Fail()
	}
}

func TestCircRemove(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.Insert(5)
	linkedList.Insert(10)
	linkedList.Remove(5)
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{6, 7, 10})
	siz := linkedList.Size() == 3
	if !seq || !siz {
		t.Fail()
	}
}

func TestCircRemoveAt(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.Insert(5)
	linkedList.Insert(10)
	linkedList.RemoveAt(1)
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{6, 5, 10})
	linkedList.RemoveAt(0)
	seq2 := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{5, 10})
	linkedList.RemoveAt(linkedList.Size() - 1)
	seq3 := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{5})
	if !seq || !seq2 || !seq3 {
		t.Fail()
	}
}

func TestCircPeekFirst(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.Insert(5)
	linkedList.Insert(10)
	seq := linkedList.PeekFirst() == 6
	if !seq {
		t.Fail()
	}
}

func TestCircPeekLast(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.Insert(5)
	linkedList.Insert(10)
	seq := linkedList.PeekLast() == 10
	if !seq {
		t.Fail()
	}
}

func TestCircPeekAt(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.Insert(5)
	linkedList.Insert(10)
	seq := linkedList.PeekAt(1) == 7
	if !seq {
		t.Fail()
	}
}

func TestCircIndexOf(t *testing.T) {
	linkedList := NewCircularLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.Insert(5)
	linkedList.Insert(10)
	seq := linkedList.IndexOf(5) == 2
	if !seq {
		t.Fail()
	}
}
