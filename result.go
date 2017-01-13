package groph

type Results struct {
	results  map[interface{}]Data
	position int
}

func NewResults() *Results {
	return &Results{
		results: make(map[interface{}]Data),
	}
}

func (r *Results) AddIfNotExists(i Data) {
	if i != nil {
		if r.results[i.GetID()] == nil {
			r.results[i.GetID()] = i
		}
	}
}

func (r *Results) GetEdge(i interface{}) *Edge {
	if r.results[i] != nil {
		return r.results[i].(*Edge)
	}

	return nil
}

func (r *Results) GetVertex(i interface{}) *Vertex {
	if r.results[i] != nil {
		return r.results[i].(*Vertex)
	}

	return nil
}

func (r *Results) Get(i interface{}) Data {
	return r.results[i]
}

func (r *Results) Vertices() (res Vertices) {
	if len(r.results) > 0 {
		for _, v := range r.results {
			if _, ok := v.(*Vertex); !ok {
				return
			}
			break
		}

		res = make([]*Vertex, 0)

		for _, v := range r.results {
			vertex := v.(*Vertex)
			res = append(res, vertex)
		}
	}

	return
}

func (r *Results) Edges() (res Edges) {
	if len(r.results) > 0 {
		for _, e := range r.results {
			if _, ok := e.(*Edge); !ok {
				return
			}
			break
		}

		res = make([]*Edge, 0)

		for _, e := range r.results {
			edge := e.(*Edge)
			res = append(res, edge)
		}
	}

	return
}
