package stack

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	s := Stack{}

	s.Push(5)
	s.Push(6)

	eq := reflect.DeepEqual(s.ToSlice(), []interface{}{5, 6})

	if !eq {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	s := Stack{}

	s.Push(5)
	s.Push(6)

	s.Pop()

	eq := reflect.DeepEqual(s.ToSlice(), []interface{}{5})

	if !eq {
		t.Fail()
	}
}

func TestPopOnEmptyStack(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	s := Stack{}

	s.Pop()
}

func TestPeek(t *testing.T) {
	s := Stack{}

	s.Push(5)
	s.Push(6)

	peekItem := s.Peek()

	eq1 := peekItem == 6
	eq2 := reflect.DeepEqual(s.ToSlice(), []interface{}{5, 6})

	if !eq1 || !eq2 {
		t.Fail()
	}

}

func TestPeekOnEmptyStack(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	s := Stack{}
	s.Peek()
}
