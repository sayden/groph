package parser

import "github.com/sayden/groph"

type edgeData struct {
	//Data   map[string]interface{} `json:"data, omitempty"`
	Weight float64 `json:"weight, omitempty"`
	//ID     string                 `json:"id"`
	groph.AnyData
}

func (e *edgeData) GetData() interface{} {
	return e.Data
}

func (e *edgeData) GetID() interface{} {
	return e.ID
}

type vertexData struct {
	//Data map[string]interface{} `json:"data, omitempty"`
	//ID   string                 `json:"id"`
	groph.AnyData
}

func (v *vertexData) GetData() interface{} {
	return v.Data
}

func (v *vertexData) GetID() interface{} {
	return v.ID
}

type ExtendedFileFormat []ExtendedRowFormat
type ExtendedRowFormat struct {
	Vertex  vertexData `json:"vertex"`
	Edge    edgeData   `json:"edge"`
	DVertex vertexData `json:"d_vertex"`
}

type SimpleFileFormat []SimpleRowFormat
type SimpleRowFormat struct {
	Vertex  string `json:"vertex"`
	Edge    string `json:"edge"`
	DVertex string `json:"d_vertex"`
}
