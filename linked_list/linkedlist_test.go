package linked_list

import (
	"reflect"
	"testing"
)

func TestInsertAtFirst(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertAtFirst(6)
	eq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{6})
	linkedList.InsertAtFirst(7)
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{7, 6})
	siz := linkedList.Size() == 2
	if !eq || !seq || !siz {
		t.Fail()
	}
}

func TestInsertAtZero(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.InsertAt(0, 5)
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{5, 6, 7})
	siz := linkedList.Size() == 3
	if !seq || !siz {
		t.Fail()
	}
}

func TestInsertAtEnd(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.InsertAt(1, 5)
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{6, 7, 5})
	siz := linkedList.Size() == 3
	if !seq || !siz {
		t.Fail()
	}
}

func TestRemoveEmptyList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Remove(5)
	eq := len(*linkedList.ToSlice())
	if eq != 0 {
		t.Fail()
	}
}

func TestRemoveOutOfIndex(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	linkedList := NewLinkedList()
	linkedList.Insert(10)
	linkedList.Insert(50)
	linkedList.RemoveAt(10)
}

func TestInsert(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(8)
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{6, 8})
	siz := linkedList.Size() == 2
	if !seq || !siz {
		t.Fail()
	}
}

func TestInsertAtLast(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.InsertAtLast(5)
	seq := reflect.DeepEqual(*linkedList.ToSlice(), []interface{}{6, 7, 5})
	siz := linkedList.Size() == 3
	if !seq || !siz {
		t.Fail()
	}
}

func TestInsertAt(t *testing.T) {
	linkedList := NewLinkedList()
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

func TestRemoveFirst(t *testing.T) {
	linkedList := NewLinkedList()
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

func TestRemoveLast(t *testing.T) {
	linkedList := NewLinkedList()
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

func TestRemove(t *testing.T) {
	linkedList := NewLinkedList()
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

func TestRemoveAt(t *testing.T) {
	linkedList := NewLinkedList()
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

func TestPeekFirst(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.Insert(5)
	linkedList.Insert(10)
	seq := linkedList.PeekFirst() == 6
	if !seq {
		t.Fail()
	}
}

func TestPeekLast(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.Insert(5)
	linkedList.Insert(10)
	seq := linkedList.PeekLast() == 10
	if !seq {
		t.Fail()
	}
}

func TestPeekAt(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.Insert(5)
	linkedList.Insert(10)
	seq := linkedList.PeekAt(1) == 7
	if !seq {
		t.Fail()
	}
}

func TestIndexOf(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Insert(6)
	linkedList.Insert(7)
	linkedList.Insert(5)
	linkedList.Insert(10)
	seq := linkedList.IndexOf(5) == 2
	if !seq {
		t.Fail()
	}
}
