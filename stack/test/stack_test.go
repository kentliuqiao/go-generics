package test

import (
	"go-generics/stack"
	"testing"
)

func Test_Push(t *testing.T) {
	s := stack.Stack[int]{}
	if s.Len() != 0 {
		t.Error("Expected 0, got", s.Len())
	}
	e, ok := s.Pop()
	if ok {
		t.Error("Expected false, got", ok)
	}
	if e != 0 {
		t.Error("Expected 0, got", e)
	}
	s.Push(2, 3)
	if s.Len() != 2 {
		t.Error("Expected 2, got", s.Len())
	}
	e, ok = s.Pop()
	if !ok {
		t.Error("Expected true, got", ok)
	}
	if e != 3 {
		t.Error("Expected 3, got", e)
	}
	if s.Len() != 1 {
		t.Error("Expected 1, got", s.Len())
	}
}
