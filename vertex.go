package groph

import "fmt"

type Vertex struct {
	Data `json:"data, omitempty"`
	OuterEdges Edges `json:"outer_edges, omitempty"`
	InnerEdges Edges `json:"inner_edges, omitempty"`
}

type Data interface {
	GetData() interface{}
	SetData(interface{})
	GetID() interface{}
	SetID(interface{})
}

// Outer returns the edges that the current vertex has
func (v *Vertex) Outer() (Edges) {
	return v.OuterEdges
}

func (v *Vertex) Inner() Edges {
	return v.InnerEdges
}

func (v *Vertex) String() string {
	res := fmt.Sprintf("VERTEX ID: '%s'\nVertex Data: %v\nEdges pointing this vertex:\n[START]", v.GetID(), v.GetData())
	for _, edge := range v.InnerEdges {
		res += edge.String()
	}

	res += "\n[END]\nEdges pointing out:\n[START]"
	for _, edge := range v.OuterEdges {
		res += edge.String()
	}

	res += "[END]"

	return res
}