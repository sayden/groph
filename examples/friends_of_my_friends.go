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

	//Inner represents edges pointing to 'v'
	v.Inner().From().Inner().From().Each(func(candidate *groph.Vertex) {

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
	})
}
