package simplestack_test

import (
	"testing"

	st "github.com/asmejia1993/simple-stack"
)

func TestEmpty(t *testing.T) {
	s := st.NewStack()
	if s.Empty() {
		t.Log("is empty")
	}
}

func TestPushToStack(t *testing.T) {
	s := st.NewStack()
	s.Push("Hello from here")
	t.Log(s.Size())
	if s.Size() > 0 {
		t.Log("Expected 1")
		t.Logf("Expected 1, found %d", s.Size())
	}
}

func TestPopFunction(t *testing.T) {
	s := st.NewStack()
	const limitToInsert = 4
	for i := 0; i < limitToInsert; i++ {
		s.Push(i + 1)
	}
	s.Pop()
	n := s.Size()
	if n == (limitToInsert - 1) {
		t.Logf("Expected %d, found %d", limitToInsert-1, n)
	}
}

func TestPeekFunction(t *testing.T) {
	s := st.NewStack()
	const limitToInsert = 4
	for i := 0; i < limitToInsert; i++ {
		s.Push(i + 1)
	}
	item, _ := s.Peek()
	if item == limitToInsert {
		t.Logf("Expected %d, found %d", item, limitToInsert)
	}
	s.Pop()
	item, _ = s.Peek()
	if item != limitToInsert {
		t.Logf("Expected 3, found %d", limitToInsert-1)
	}
}
