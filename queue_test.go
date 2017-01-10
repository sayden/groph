package groph

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue()

	_, err := q.Pop()
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	q.Push(&Edge{Weight:2, Data:&AnyData{Data:"hello", ID:"hello"}})

	v, err := q.Pop()
	if err != nil {
		t.Fail()
	}

	if v.Data.GetID() != "hello" || v.Weight != 2 {
		t.Fail()
	}
}
