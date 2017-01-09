package groph

import "fmt"

type Vertex struct {
	Data `json:"data, omitempty"`
	OuterEdges []*Edge `json:"outer_edges, omitempty"`
	InnerEdges []*Edge `json:"inner_edges, omitempty"`
}

type Data interface {
	GetData() interface{}
	GetID() interface{}
}

// Outer returns the edges that the current vertex has
func (v *Vertex) Outer() ([]*Edge, error) {
	if len(v.OuterEdges) == 0 {
		return nil, noEdgesError
	}

	return v.OuterEdges, nil
}

func (v *Vertex) Inner() ([]*Edge, error) {
	if len(v.InnerEdges) == 0 {
		return nil, noEdgesError
	}

	return v.InnerEdges, nil
}

func (v *Vertex) String() string {
	res := fmt.Sprintf("Data: %s(%s)\n", v.Data.GetData(), v.Data.GetID())
	for _, edge := range v.InnerEdges {
		res += edge.String()
	}

	for _, edge := range v.OuterEdges {
		res += edge.String()
	}

	return res
}
