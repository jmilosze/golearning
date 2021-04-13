package datastructs

import (
	"reflect"
	"testing"
)

func TestGraph_AddNode(t *testing.T) {
	tests := []struct {
		name         string
		graph        Graph
		newNodeValue int
		want         *Graph
		wantErr      bool
	}{
		{
			"adding existing node causes error",
			Graph{10: make(map[int]int)},
			10,
			&Graph{10: make(map[int]int)},
			true,
		},
		{
			"adding new node adds node",
			Graph{10: make(map[int]int)},
			20,
			&Graph{10: make(map[int]int), 20: make(map[int]int)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.graph.AddNode(tt.newNodeValue)
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

func TestGraph_AddConnection(t *testing.T) {
	type args struct {
		valueFrom int
		valueTo   int
		weight    int
	}
	tests := []struct {
		name     string
		graph    Graph
		connArgs args
		want     *Graph
		wantErr  bool
	}{
		{
			"adding connection from non-existing node causes error",
			Graph{10: make(map[int]int), 20: make(map[int]int), 30: make(map[int]int)},
			args{5, 20, 7},
			&Graph{10: make(map[int]int), 20: make(map[int]int), 30: make(map[int]int)},
			true,
		},
		{
			"adding connection to non-existing node causes error",
			Graph{10: make(map[int]int), 20: make(map[int]int), 30: make(map[int]int)},
			args{10, 5, 7},
			&Graph{10: make(map[int]int), 20: make(map[int]int), 30: make(map[int]int)},
			true,
		},
		{
			"adding connection adds connection correctly",
			Graph{10: make(map[int]int), 20: make(map[int]int), 30: make(map[int]int)},
			args{10, 20, 7},
			&Graph{10: map[int]int{20: 7}, 20: make(map[int]int), 30: make(map[int]int)},
			false,
		},
		{
			"adding connection overrides existing connection",
			Graph{10: map[int]int{20: 7}, 20: make(map[int]int), 30: make(map[int]int)},
			args{10, 20, 1},
			&Graph{10: map[int]int{20: 1}, 20: make(map[int]int), 30: make(map[int]int)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.graph.AddConnection(tt.connArgs.valueFrom, tt.connArgs.valueTo, tt.connArgs.weight)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddConnection() got = %v, want %v", got, tt.want)
			}
		})
	}
}
