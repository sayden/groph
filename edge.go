package groph

import "fmt"

type Edge struct {
	Data `json:"data, omitempty"`
	PointsTo *Vertex     `json:"PointsTo, omitempty"`
	From     *Vertex     `json:"From, omitempty"`
	Weight   float64     `json:"weight, omitempty"`
}

func (e *Edge) Points(v *Vertex) bool {
	return e.PointsTo == v
}

func (e *Edge) String() string {
	return fmt.Sprintf("\nEDGE ID: '%s'\nData: '%v'\nPointsTo ID: '%s'\nPointsTo Data: '%v'\nFrom ID: '%s'\nFrom Data: '%v'\nWeight: '%f'\n",
		e.GetID(), e.GetData(), e.PointsTo.GetID(), e.PointsTo.GetData(), e.From.GetID(), e.From.GetData(), e.Weight)
}
