package groph

import "testing"

func TestResults_Vertices(t *testing.T) {
	res := Results{
		results:make(map[interface{}]Data),
	}

	graph := New()

	any1 := graph.NewVertex(&AnyData{Data:"data",ID:"id1"})
	any2 := graph.NewVertex(&AnyData{Data:"data",ID:"id2"})

	res.AddIfNotExists(any1)
	res.AddIfNotExists(any2)

	final := res.Vertices()

	if final == nil {
		t.Fatal()
	}

	final.Each(func(v *Vertex){
		if v.GetID() != "id1" && v.GetID() != "id2" {
			t.Fail()
		}
	})

	d := res.Get("id1")
	if d == nil {
		t.Fail()
	}

	if d.GetID() != "id1" {
		t.Fail()
	}

	if res.GetVertex("id1") == nil {
		t.Fail()
	}

	if res.GetVertex("id1").GetID() != "id1"{
		t.Fail()
	}
}

func TestResults_Edges(t *testing.T) {
	res := Results{
		results:make(map[interface{}]Data),
	}

	graph := New()

	any1 := graph.NewEdge(&AnyData{Data:"data",ID:"id1"}, 0)
	any2 := graph.NewEdge(&AnyData{Data:"data",ID:"id2"}, 0)

	res.AddIfNotExists(any1)
	res.AddIfNotExists(any2)

	final := res.Edges()

	if final == nil {
		t.Fatal()
	}

	final.Each(func(v *Edge){
		if v.GetID() != "id1" && v.GetID() != "id2" {
			t.Fail()
		}
	})

	d := res.Get("id1")
	if d == nil {
		t.Fail()
	}

	if d.GetID() != "id1" {
		t.Fail()
	}

	if res.GetEdge("id1") == nil {
		t.Fail()
	}

	if res.GetEdge("id1").GetID() != "id1"{
		t.Fail()
	}
}

func TestResults_AddIfNotExists(t *testing.T) {
	res := Results{
		results:make(map[interface{}]Data),
	}

	any := &AnyData{
		Data:"data",
		ID:"id",
	}

	res.results[any.ID] = any

	res.AddIfNotExists(any)

	count := 0
	for range res.results {
		count++
	}

	if count != 1 {
		t.Fail()
	}
}

