package groph

import (
	"math"
	"fmt"
)

func (g *Graph) ShortestPathWithVertex(t *Vertex) (reversed []*Vertex, totalCost float64, err error) {
	reversed, totalCost, err = g.ShortestPath(t.GetID())
	reversed = append(reversed, t)
	return
}

func (g *Graph) ShortestPath(id interface{}) (reversed Vertices, totalCost float64, err error) {
	costs := make(map[interface{}]float64)
	costs[id] = math.Inf(1)

	parents := make(map[interface{}]*Edge)
	parents[id] = nil

	queue := NewQueue()
	for _, e := range g.StartVertex.OuterEdges {
		queue.Push(e)
		costs[e.PointsTo.GetID()] = math.Inf(1)
		parents[e.PointsTo.GetID()] = e
	}

	err = g.shortest(id, costs, queue, parents, make(map[interface{}]bool))
	if err != nil {
		return
	}

	//Traverse back
	res := []*Vertex{}
	cur := parents[id]

	for cur.From != g.StartVertex {
		totalCost += cur.Weight
		res = append(res, cur.From)
		if parents[cur.From.GetID()] == nil {
			return nil, 0, fmt.Errorf("No parent found for %s\n", cur.From.GetID())
		}
		cur = parents[cur.From.GetID()]
	}
	res = append(res, g.StartVertex)
	totalCost += cur.Weight

	reversed = make([]*Vertex, len(res))

	var j int
	for i := len(res) - 1; i >= 0; i-- {
		reversed[j] = res[i]
		j++
	}

	reversed = append(reversed, parents[id].PointsTo)

	return
}

func (g *Graph) shortest(id interface{}, costs map[interface{}]float64, queue *Queue, parents map[interface{}]*Edge, processed map[interface{}]bool) (err error) {

	currentEdge, err := queue.Pop()
	if err != nil {
		return nil
	}

	//Don't process outer edges of target node
	if currentEdge.From.GetID() == id {
		return
	}

	if !processed[currentEdge] {

		if costs[currentEdge.PointsTo.GetID()] > currentEdge.Weight {
			costs[currentEdge.PointsTo.GetID()] = currentEdge.Weight
			parents[currentEdge.PointsTo.GetID()] = currentEdge
		}

		for _, edge := range currentEdge.PointsTo.OuterEdges {
			queue.Push(edge)
		}

		processed[currentEdge] = true
	}

	return g.shortest(id, costs, queue, parents, processed)
}
