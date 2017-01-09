# groph
A WIP simple database library written in Go

This is more like a research project around an idea I had of making queries on a simple using callbacks than a serious simple database ready for production use.

You can load data in the graph by usin JSON or YAML data in "triples" format or a more advanced "extended" format:

## Downloading and installing

```bash
go get -u github.com/sayden/groph
```

## Quickstart
TODO

## Simple format

An example YAML simple format file:

```yaml
- vertex: start
  edge: goes to
  d_vertex: B

- vertex: A
  edge: goes to
  d_vertex: finish

- vertex: A
  edge: goes to
  d_vertex: B

- vertex: B
  edge: goes to
  d_vertex: finish

- vertex: B
  edge: goes to
  d_vertex: A

- vertex: finish
  edge: goes to
  d_vertex: start
```

As you can see you just have to declare the source vertex, the destination vertex and an ID for the edge that connects them.

The JSON equivalent file would be like this:

```json
[
  { "vertex": "start", "edge": "goes to", "d_vertex": "B" },
  { "vertex": "A", "edge": "goes to", "d_vertex": "finish" },
  { "vertex": "A", "edge": "goes to", "d_vertex": "B" },
  { "vertex": "B", "edge": "goes to", "d_vertex": "finish" },
  { "vertex": "B", "edge": "goes to", "d_vertex": "A" },
  { "vertex": "finish", "edge": "goes to", "d_vertex": "start" }
]
````

## Extended format

In the case of the extended format:

```yaml
- vertex:
    id: start
    data:
      description: starting point
  edge:
    id: goes to
    data:
      description: data in some edge
    weight: 52.3
  d_vertex:
    id: A

- vertex:
    id: A
  edge:
    id: goes to
    weight: 52.3
  d_vertex:
    id: finish

- vertex:
    id: A
  edge:
    id: goes to
  d_vertex:
    id: B

- vertex:
    id: B
  edge:
    id: goes to
  d_vertex:
    id: A

- vertex:
    id: B
  edge:
    id: goes to
    weight: 52.3
  d_vertex:
    id: finish

```

In the extended format each vertex and edge can contain a "data" key with any value that you can consider inside. In this case the ID field is mandatory on each node of the YAML / JSON.


And the JSON equivalent:

```json
[
  {
    "vertex":  { "id": "start",   "data":{ "description":"starting point"    }                },
    "edge":    { "id": "goes to", "data":{ "description":"data in some edge" }, "weight":52.3 },
    "d_vertex":{ "id": "A"                                                                    }
  },
  {
    "vertex":  { "id": "A" },
    "edge":    { "id": "goes to", "weight":52.3 },
    "d_vertex":{ "id": "finish"}
  },
  {
    "vertex":  { "id": "A" },
    "edge":    { "id": "goes to" },
    "d_vertex":{ "id": "B"}
  },
  {
    "vertex":  { "id": "B" },
    "edge":    { "id": "goes to" },
    "d_vertex":{ "id": "A"}
  },
  {
    "vertex":  { "id": "B" },
    "edge":    { "id": "goes to", "weight":52.3 },
    "d_vertex":{ "id": "finish"}
  }
]
````

If you want a rock-solid simple database written in Go check [Cayley](https://github.com/cayleygraph/cayley)