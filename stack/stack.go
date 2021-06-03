package stack

type Stack struct {
	items []interface{}
	top   int
}

func (s *Stack) IsEmpty() bool {
	return len((*s).items) == 0
}

func (s *Stack) Push(a interface{}) {
	(*s).items = append((*s).items, a)
	(*s).top++
}

func (s *Stack) Peek() interface{} {
	return (*s).items[(*s).top-1]
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		panic("pop on empty stack is not allowed")
	}
	subArr := (*s).items[:(*s).top-1]
	newArr := make([]interface{}, len(subArr))
	copy(newArr, subArr)
	
	(*s).items = newArr
}

func (s *Stack) ToSlice() []interface{} {
	return (*s).items
}
