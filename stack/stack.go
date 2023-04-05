package stack

type Stack[E comparable] []E

func (s *Stack[E]) Pop() (E, bool) {
	var e E
	if len(*s) == 0 {
		return e, false
	}
	e = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return e, true
}

func (s *Stack[E]) Push(es ...E) {
	*s = append(*s, es...)
}

func (s *Stack[E]) Len() int {
	return len(*s)
}
