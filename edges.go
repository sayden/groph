package groph

type Edges []*Edge

func (es Edges) Map(f func(*Edge) *Edge) (res Edges) {
	res = make([]*Edge, len(es))

	for i:=0;i<len(es);i++{
		res[i] = f(es[i])
	}

	return
}

func (es Edges) Each(f func(*Edge)) {
	for i:=0;i<len(es);i++{
		f(es[i])
	}
}

func (es Edges) Filter(f func(*Edge) bool) (res Edges) {
	res = make([]*Edge, 0)

	for i:=0;i<len(es);i++{
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