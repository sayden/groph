package main

import (
	"fmt"

	"github.com/sayden/groph"
	"github.com/sayden/groph/parser"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

//Nodes that points to nodes that points to 'finish'
func main() {
	graph, err := parser.ReadJSONExtendedFileFormat("parser/test.json")
	panicIfErr(err)

	v, err := graph.Find("finish")
	panicIfErr(err)

	// From previous result, tell me shortes routes from 'finish' to any of vertices in result if exists
	graph.SetRootVertex(v)
	noDuplicateSearch := make(map[interface{}]bool)

	//Inner represents edges pointing to 'v'
	v.Inner().MapV(func(edge *groph.Edge) *groph.Vertex{
		return edge.From

	}).FlatMap(func(v *groph.Vertex) groph.Vertices {
		//Each edge has a 'From' and a 'PointTo' field with the references to the vertices it connects
		return v.Inner().MapV(func(e *groph.Edge) *groph.Vertex{
			return e.From
		})

	}).Each(func(candidate *groph.Vertex) {
		if !noDuplicateSearch[candidate.GetID()] {
			noDuplicateSearch[candidate.GetID()] = true

			fmt.Printf("Searching shortest route from '%s' to '%s'\n", v.GetID(), candidate.GetID())
			result, totalCost, err := graph.ShortestPath(candidate.GetID())

			if err != nil {
				fmt.Printf("Could not find route: %s\n", err)

			} else {
				fmt.Printf("Path found from finish to %s with a cost of %f\n", candidate.GetID(), totalCost)

				result.Each(func(v *groph.Vertex) {
					fmt.Printf("%s -> ", v.GetID())
				})

				fmt.Printf("\n\n")
			}
		}
	})
}
