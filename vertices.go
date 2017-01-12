package groph

type Vertices []*Vertex

type VerticesSlice []Vertices

func (vs Vertices) Map(f func(v *Vertex) *Vertex) (res Vertices){
	res = make([]*Vertex, len(vs))

	for k, vertex := range vs {
		res[k] = f(vertex)
	}

	return
}

func (vs Vertices) FlatMap(f func(v *Vertex) Vertices) (res Vertices){
	res = make([]*Vertex, 0)

	for _, vertices := range vs {
		res = append(res, f(vertices)...)
	}

	return
}


func (vs Vertices) MapE(f func(v *Vertex) *Edge) (res Edges){
	res = make([]*Edge, len(vs))

	for k, vertex := range vs {
		res[k] = f(vertex)
	}

	return
}


func (vs Vertices) MapT(f func(v *Vertex) interface{}) (res []interface{}){
	res = make([]interface{}, len(vs))

	for k, vertex := range vs {
		res[k] = f(vertex)
	}

	return
}

func (vs Vertices) Each(f func(v *Vertex)){
	for _, vertex := range vs {
		f(vertex)
	}
}

func (vs Vertices) Filter(f func(v *Vertex) bool) (res Vertices){
	res = make([]*Vertex, 0)

	for _, vertex := range vs {
		if f(vertex) {
			res = append(res, vertex)
		}
	}

	return
}

func (vs Vertices) Fold(init interface{}, f func(a interface{}, b *Vertex) interface{}) (cur interface{}) {
	cur = init

	for _, vertex := range vs {
		cur = f(cur, vertex)
	}

	return cur
}
