package parser

import (
	"os"

	"encoding/json"

	"io/ioutil"

	"github.com/sayden/groph"
	"github.com/ghodss/yaml"
)

// ReadJSONFile will try to return a parsed ExtendedFileFormat from a JSON file. The extended file format is a special
// JSON format to load graph data
func readJSONFileFormat(filePath string, fileType interface{}) (err error){
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		return
	}

	err = json.NewDecoder(file).Decode(fileType)
	return
}

// ReadYAMLFile will try to return a parsed ExtendedFileFormat form a YAML file. The same rules that JSON applies to
// this file
func readYAMLFileFormat(filePath string, graphFile interface{}) (err error) {
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		return
	}

	var yml []byte
	yml, err = ioutil.ReadAll(file)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(yml, &graphFile)

	return
}
// ConvertSimpleFileFormat
func ConvertSimpleFileFormat(gf *SimpleFileFormat) (graph *groph.Graph, err error) {
	graph = groph.NewGraph()

	for _, row := range *gf {
		addSimpleRow(row, graph)
	}

	return
}

// ConvertExtended must be called to create a graph from a ExtendedFileFormat value.
func ConvertExtended(gf *ExtendedFileFormat) (graph *groph.Graph, err error) {
	graph = groph.NewGraph()

	for _, row := range *gf {
		addExtendedRow(row, graph)
	}

	return
}

func addSimpleRow(r SimpleRowFormat, g *groph.Graph) {
	sourceData := &groph.AnyData{
		Data:r.Vertex,
		ID:r.Vertex,
	}
	source := g.NewVertex(sourceData)

	destData := &groph.AnyData{
		Data:r.DVertex,
		ID:r.DVertex,
	}
	dest := g.NewVertex(destData)

	edgeData := &groph.AnyData{
		Data:r.Edge,
		ID:r.Edge,
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
