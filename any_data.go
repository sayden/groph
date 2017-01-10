package groph

// AnyData is a generic type to hold data within a vertex or an edge. One of its uses in the library is to parse
// unknown incoming data from a JSON or YAML file
type AnyData struct {
	Data interface{} `json:"data, omitempty"`
	ID   interface{} `json:"id, omitempty"`
}

func (a *AnyData) GetData() interface{} {
	return a.Data
}

func (a *AnyData) GetID() interface{} {
	return a.ID
}

func (a *AnyData) SetID(i interface{}){
	a.ID = i
}

func (a *AnyData) SetData(d interface{}){
	a.Data = d
}