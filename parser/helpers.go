package parser

import (
	"encoding/json"

	"github.com/sayden/groph"
)

// convertSimpleFileFormat
func convertSimpleFileFormat(gf *SimpleFileFormat) (graph *groph.Graph, err error) {
	graph = groph.New()

	for _, row := range *gf {
		addSimpleRow(row, graph)
	}

	return
}

// ConvertExtendedFormat must be called to create a graph from a ExtendedFileFormat value.
func ConvertExtendedFormat(gf *ExtendedFileFormat) (graph *groph.Graph, err error) {
	graph = groph.New()

	for _, row := range *gf {
		addExtendedRow(row, graph)
	}

	return
}

// ParseExtendedFormatBytes can be used to parse an array of bytes to a graph if you aren't loading the graph from a file
func ParseExtendedFormatBytes(b []byte) (graph *groph.Graph, err error) {
	extended := new(ExtendedFileFormat)
	err = json.Unmarshal(b, &extended)
	if err != nil {
		return
	}

	graph, err = ConvertExtendedFormat(extended)

	return
}

// ParseSimpleFormatBytes can be used to parse an array of bytes to a graph if you aren't loading the graph from a file
func ParseSimpleFormatBytes(b []byte) (graph *groph.Graph, err error) {
	simple := new(SimpleFileFormat)
	err = json.Unmarshal(b, &simple)
	if err != nil {
		return
	}

	graph, err = convertSimpleFileFormat(simple)

	return
}

func addSimpleRow(r SimpleRowFormat, g *groph.Graph) {
	sourceData := &groph.AnyData{
		Data: r.Vertex,
		ID:   r.Vertex,
	}
	source := g.NewVertex(sourceData)

	destData := &groph.AnyData{
		Data: r.DVertex,
		ID:   r.DVertex,
	}
	dest := g.NewVertex(destData)

	edgeData := &groph.AnyData{
		Data: r.Edge,
		ID:   r.Edge,
	}
	edge := g.NewEdge(edgeData, 0)

	g.AddConnection(source, dest, edge)
}

func addExtendedRow(r ExtendedRowFormat, g *groph.Graph) {
	source := g.NewVertex(&r.Vertex)
	dest := g.NewVertex(&r.DVertex)

	edge := g.NewEdge(&r.Edge, r.Edge.Weight)

	g.AddConnection(source, dest, edge)
}
