package groph

type Edges []*Edge

func (es Edges) Map(f func(*Edge) *Edge) (res Edges) {
	res = make([]*Edge, len(es))

	for i := 0; i < len(es); i++ {
		res[i] = f(es[i])
	}

	return
}

func (es Edges) MapV(f func(*Edge) *Vertex) (res Vertices) {
	res = make([]*Vertex, len(es))

	for i := 0; i < len(es); i++ {
		res[i] = f(es[i])
	}

	return
}

func (es Edges) Size() int {
	return len(es)
}

func (es Edges) MapT(f func(*Edge) Data) (res *Results) {
	res = NewResults()

	for i := 0; i < len(es); i++ {
		res.AddIfNotExists(f(es[i]))
	}

	return
}

func (es Edges) Each(f func(*Edge)) {
	for i := 0; i < len(es); i++ {
		f(es[i])
	}
}

func (es Edges) Filter(f func(*Edge) bool) (res Edges) {
	res = make([]*Edge, 0)

	for i := 0; i < len(es); i++ {
		if f(es[i]) {
			res = append(res, es[i])
		}
	}

	return
}

func (es Edges) Fold(init interface{}, f func(a interface{}, b *Edge) interface{}) (cur interface{}) {
	cur = init

	for _, edge := range es {
		cur = f(cur, edge)
	}

	return cur
}

func (es Edges) Outer() Vertices {
	res := NewResults()

	for _, v := range es {
		res.AddIfNotExists(v.PointsTo)
	}

	return res.Vertices()
}

func (es Edges) Inner() Vertices {
	res := NewResults()

	for _, v := range es {
		res.AddIfNotExists(v.From)
	}

	return res.Vertices()
}