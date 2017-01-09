package groph

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue()

	q.Push(&Edge{Weight:2, Data:&EdgeMockData{Data:"hello"}})

	v, err := q.Pop()
	if err != nil {
		t.Fail()
	}

	if v.Data.GetID() != "hello" || v.Weight != 2 {
		t.Fail()
	}
}
