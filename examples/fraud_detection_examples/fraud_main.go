package main

import (
	"fmt"

	"github.com/sayden/groph"
	"github.com/sayden/groph/parser"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

type item struct {
	Item string `json:"type"`
}

//Nodes that points to nodes that points to 'finish'
func main() {
	graph, err := parser.ReadJSONExtendedFileFormat("examples/fraud_detection_examples/fraud_data.json")
	panicIfErr(err)

	//Detects users with more than 2 payment methods (representing N payment methods that could be fake)
	graph.Filter(func(v *groph.Vertex) bool {
		data, _ := v.GetData().(map[string]interface{})
		return data["type"] == "User"
	}).Each(func(v *groph.Vertex) {
		pmNumber := v.InEdges().Filter(func(e *groph.Edge) bool {
			s, _ := e.GetID().(string)
			return s == "from_user"
		}).Size()

		if pmNumber > 2 {
			fmt.Printf("Potential fraudster: %s has %d payment methods\n", v.GetID(), pmNumber)
		}
	})

	//Users sharing a payment method (representing a fraud ring)
	graph.Filter(func(v *groph.Vertex) bool {
		data, _ := v.GetData().(map[string]interface{})
		return data["type"] == "PaymentMethod"
	}).Each(func(v *groph.Vertex) {
		users := make(map[interface{}][]interface{})

		v.OutEdges().Filter(func(e *groph.Edge) bool {
			return fmt.Sprintf("%s", e.GetID()) == "from_user"
		}).Each(func(e *groph.Edge) {
			if users[e.From.GetID()] == nil {
				users[e.From.GetID()] = make([]interface{}, 0)
			}
			users[e.From.GetID()] = append(users[e.From.GetID()], e.PointsTo.GetID())

			if len(users[e.From.GetID()]) > 1 {
				fmt.Printf("Payment Method %s is shared by the following users:\n", e.From.GetID())
				for _, v := range users[e.From.GetID()] {
					fmt.Println(v)
				}
			}
		})
	})
}
