package parser

import (
	"os"

	"encoding/json"

	"github.com/sayden/groph"
)

// ReadJSONSimpleFileFormat is a helper function that will take a file path with a JSON file in simple format and
// return a graph value
func ReadJSONSimpleFileFormat(filePath string) (graph *groph.Graph, err error) {
	simple := new(SimpleFileFormat)
	err = readJSONFileFormat(filePath, simple)
	if err != nil {
		return
	}

	graph, err = convertSimpleFileFormat(simple)

	return
}

// ReadJSONExtendedFileFormat is a helper function that will take a file path with a JSON file in extended format and
// it will return a graph value
func ReadJSONExtendedFileFormat(filePath string) (graph *groph.Graph, err error) {
	simple := new(ExtendedFileFormat)
	err = readJSONFileFormat(filePath, simple)
	if err != nil {
		return
	}

	graph, err = ConvertExtendedFormat(simple)

	return
}

func readJSONFileFormat(filePath string, fileType interface{}) (err error) {
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		return
	}

	err = json.NewDecoder(file).Decode(fileType)
	return
}
