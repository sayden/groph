package groph

import "errors"

type Queue []*Edge

func (q *Queue) Push(v *Edge) {
	(*q) = append(*q, v)
}

func (q *Queue) Pop() (ret *Edge, err error) {
	if len((*q)) > 0 {
		ret = (*q)[0]
		(*q) = (*q)[1:]
		return
	}

	return nil, errors.New("Not enough elements")
}

func NewQueue() *Queue {
	var q Queue
	q = make([]*Edge, 0)
	return &q
}
