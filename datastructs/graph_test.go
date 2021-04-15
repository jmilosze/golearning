package datastructs

import (
	"reflect"
	"testing"
)

func TestGraph_AddNode(t *testing.T) {
	node := GNode{make(map[int]int), 0}

	tests := []struct {
		name         string
		graph        Graph
		newNodeValue int
		want         *Graph
		wantErr      bool
	}{
		{
			"adding existing node causes error",
			Graph{10: node},
			10,
			&Graph{10: node},
			true,
		},
		{
			"adding node smaller than 1 node causes error",
			Graph{10: node},
			0,
			&Graph{10: node},
			true,
		},
		{
			"adding new node adds node",
			Graph{10: node},
			20,
			&Graph{10: node, 20: node},
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
	nodeNoConn := GNode{make(map[int]int), 0}

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
			"adding connection from non-existing nodeNoConn causes error",
			Graph{10: nodeNoConn, 20: nodeNoConn, 30: nodeNoConn},
			args{5, 20, 7},
			&Graph{10: nodeNoConn, 20: nodeNoConn, 30: nodeNoConn},
			true,
		},
		{
			"adding connection to non-existing nodeNoConn causes error",
			Graph{10: nodeNoConn, 20: nodeNoConn, 30: nodeNoConn},
			args{10, 5, 7},
			&Graph{10: nodeNoConn, 20: nodeNoConn, 30: nodeNoConn},
			true,
		},
		{
			"adding connection to itself causes error",
			Graph{10: nodeNoConn, 20: nodeNoConn, 30: nodeNoConn},
			args{10, 10, 1},
			&Graph{10: nodeNoConn, 20: nodeNoConn, 30: nodeNoConn},
			true,
		},
		{
			"adding connection with negative weight causes error",
			Graph{10: nodeNoConn, 20: nodeNoConn, 30: nodeNoConn},
			args{10, 20, -1},
			&Graph{10: nodeNoConn, 20: nodeNoConn, 30: nodeNoConn},
			true,
		},
		{
			"adding connection adds connection correctly",
			Graph{10: nodeNoConn, 20: nodeNoConn, 30: nodeNoConn},
			args{10, 20, 7},
			&Graph{10: GNode{map[int]int{20: 7}, 0}, 20: nodeNoConn, 30: nodeNoConn},
			false,
		},
		{
			"adding connection overrides existing connection",
			Graph{10: GNode{map[int]int{20: 7}, 0}, 20: nodeNoConn, 30: nodeNoConn},
			args{10, 20, 1},
			&Graph{10: GNode{map[int]int{20: 1}, 0}, 20: nodeNoConn, 30: nodeNoConn},
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
