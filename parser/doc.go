// Package parser provides function to read JSON and YAML files and convert them to a graph value.
//
// To load the graph, you can use two type of syntax: simple and extended.
//
// Simple:
//
// The simple syntax is the best to get started with the library. You just have to define as many triples as you need
// https://en.wikipedia.org/wiki/Semantic_triple in JSON or YAML format. For example in JSON:
//
//  [
//	    { "vertex": "start", "edge": "goes to", "d_vertex": "B" },
//	    { "vertex": "A", "edge": "goes to", "d_vertex": "finish" },
//	    { "vertex": "A", "edge": "goes to", "d_vertex": "B" },
//	    { "vertex": "B", "edge": "goes to", "d_vertex": "finish" },
//	    { "vertex": "B", "edge": "goes to", "d_vertex": "A" },
//	    { "vertex": "finish", "edge": "goes to", "d_vertex": "start" }
//  ]
//
// Or in YAML:
//
//		- vertex: start
//		  edge: goes to
//		  d_vertex: B
//
//		- vertex: A
//		  edge: goes to
//		  d_vertex: finish
//
//		- vertex: A
//		  edge: goes to
//		  d_vertex: B
//
//		- vertex: B
//		  edge: goes to
//		  d_vertex: finish
//
//		- vertex: B
//		  edge: goes to
//		  d_vertex: A
//
//		- vertex: finish
//		  edge: goes to
//		  d_vertex: start
//
//
// Once you have a file, you can load it with the helper functions:
//          ReadJSONSimpleFileFormat(filePath string)(graph *groph.Graph, err error)
//          ReadJSONExtendedFileFormat(filePath string) (graph *groph.Graph, err error)
//          ReadYAMLSimpleFileFormat(filePath string) (graph *groph.Graph, err error)
//          ReadYAMLExtendedFileFormat(filePath string)(graph *groph.Graph, err error)
//
// The library will recognize vertex, edge, and d_vertex names as ID's so, for example, a duplicate name on any vertex
// will create a new edge between vertices. And edge must always been supplied with a valid ID.
//
// Extended:
//
// The extended file format can contain lot of information that will be stored within the edges and vertices. An
// example of this format in JSON is like this:
//
//		[
//			{
//				"vertex":  { "id": "start",   "data":{ "description":"starting point"    }                },
//				"edge":    { "id": "goes to", "data":{ "description":"data in some edge" }, "weight":52.3 },
//				"d_vertex":{ "id": "A"                                                                    }
//			},
//			{
//				"vertex":  { "id": "A" },
//				"edge":    { "id": "goes to", "weight":52.3 },
//				"d_vertex":{ "id": "finish"}
//			},
//			{
//				"vertex":  { "id": "A" },
//				"edge":    { "id": "goes to" },
//				"d_vertex":{ "id": "B"}
//			},
//			{
//				"vertex":  { "id": "B" },
//				"edge":    { "id": "goes to" },
//				"d_vertex":{ "id": "A"}
//			},
//			{
//				"vertex":  { "id": "B" },
//				"edge":    { "id": "goes to", "weight":52.3 },
//				"d_vertex":{ "id": "finish"}
//			}
//		]
//
// Or in YAML:
//      - vertex:
//          id: start
//          data:
//            description: starting point
//        edge:
//          id: goes to
//          data:
//            description: data in some edge
//          weight: 52.3
//        d_vertex:
//          id: A
//
//      - vertex:
//          id: A
//        edge:
//          id: goes to
//          weight: 52.3
//        d_vertex:
//          id: finish
//
//      - vertex:
//          id: A
//        edge:
//          id: goes to
//        d_vertex:
//          id: B
//
//      - vertex:package parser
//          id: B
//        edge:
//          id: goes to
//        d_vertex:
//          id: A
//
//      - vertex:
//          id: B
//        edge:
//          id: goes to
//          weight: 52.3
//        d_vertex:
//          id: finish
package parser