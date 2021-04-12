package datastructs

import (
	"reflect"
	"testing"
)

func TestGraph_AddNode(t *testing.T) {
	tests := []struct {
		name         string
		initGraph    Graph
		newNodeValue int
		want         *Graph
		wantErr      bool
	}{
		{
			"adding existing node causes error",
			Graph{10: make(map[int]int)},
			10,
			&(Graph{10: make(map[int]int)}),
			true,
		},
		{
			"adding new node adds node",
			Graph{10: make(map[int]int)},
			20,
			&(Graph{10: make(map[int]int), 20: make(map[int]int)}),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.initGraph.AddNode(tt.newNodeValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
