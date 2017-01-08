package groph

import "fmt"

type Edge struct {
	Data     interface{} `json:"data, omitempty"`
	PointsTo *Vertex     `json:"PointsTo, omitempty"`
	From     *Vertex     `json:"From, omitempty"`
	Weight   float64     `json:"weight, omitempty"`
}

func (e *Edge) Points(v *Vertex) bool {
	return e.PointsTo == v
}

func (e *Edge) String() string {
	return fmt.Sprintf("\n%#v\nPointsTo: %s(%s)\nFrom: %s(%s)\nWeight: %f\n", e.Data, e.PointsTo.GetData(), e.PointsTo.GetID(), e.From.GetData(), e.From.GetID(), e.Weight)
}
