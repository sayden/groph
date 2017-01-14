package groph

import "fmt"

type Vertex struct {
	Data     `json:"data, omitempty"`
	outEdges Edges
	inEdges  Edges
}

type Data interface {
	GetData() interface{}
	SetData(interface{})
	GetID() interface{}
	SetID(interface{})
}

// OutEdges returns the edges of the pointing vertex
func (v *Vertex) OutEdges() Edges {
	return v.outEdges
}

func (v *Vertex) InEdges() Edges {
	return v.inEdges
}

func (v *Vertex) String() string {
	res := fmt.Sprintf("VERTEX ID: '%s'\nVertex Data: %v\nEdges pointing this vertex:\n[START]", v.GetID(), v.GetData())
	for _, edge := range v.inEdges {
		res += edge.String()
	}

	res += "\n[END]\nEdges pointing out:\n[START]"
	for _, edge := range v.outEdges {
		res += edge.String()
	}

	res += "[END]"

	return res
}

func PrintV(v *Vertex) {
	fmt.Println(v)
}
