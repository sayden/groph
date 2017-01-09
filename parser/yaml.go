package parser

import (
	"io/ioutil"
	"os"

	"github.com/sayden/groph"
	"github.com/ghodss/yaml"
)

// ReadYAMLSimpleFileFormat is a helper function that will take a file path with a YAML file in simple format and
// return a graph value
func ReadYAMLSimpleFileFormat(filePath string) (graph *groph.Graph, err error) {
	simple := new(SimpleFileFormat)
	err = readYAMLFileFormat(filePath, simple)
	if err != nil {
		return
	}

	graph, err = convertSimpleFileFormat(simple)

	return
}

// ReadYAMLExtendedFileFormat is a helper function that will take a file path with a YAML file in extended format and
// return a graph value
func ReadYAMLExtendedFileFormat(filePath string)(graph *groph.Graph, err error){
	extended := new(ExtendedFileFormat)
	err = readYAMLFileFormat(filePath, extended)
	if err != nil {
		return
	}

	graph, err = ConvertExtendedFormat(extended)

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
