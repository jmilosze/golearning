package datastructs

import (
	"reflect"
	"testing"
)

func emptyNode() *GNode {
	return &GNode{make(map[int]int)}
}

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
			Graph{10: emptyNode()},
			10,
			&Graph{10: &GNode{make(map[int]int)}},
			true,
		},
		{
			"adding node smaller than 1 node causes error",
			Graph{10: emptyNode()},
			0,
			&Graph{10: emptyNode()},
			true,
		},
		{
			"adding new node adds node",
			Graph{10: emptyNode()},
			20,
			&Graph{10: emptyNode(), 20: emptyNode()},
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
			"adding connection from non-existing nodeNoConn causes error",
			Graph{10: emptyNode(), 20: emptyNode(), 30: emptyNode()},
			args{5, 20, 7},
			&Graph{10: emptyNode(), 20: emptyNode(), 30: emptyNode()},
			true,
		},
		{
			"adding connection to non-existing nodeNoConn causes error",
			Graph{10: emptyNode(), 20: emptyNode(), 30: emptyNode()},
			args{10, 5, 7},
			&Graph{10: emptyNode(), 20: emptyNode(), 30: emptyNode()},
			true,
		},
		{
			"adding connection to itself causes error",
			Graph{10: emptyNode(), 20: emptyNode(), 30: emptyNode()},
			args{10, 10, 1},
			&Graph{10: emptyNode(), 20: emptyNode(), 30: emptyNode()},
			true,
		},
		{
			"adding connection with negative weight causes error",
			Graph{10: emptyNode(), 20: emptyNode(), 30: emptyNode()},
			args{10, 20, -1},
			&Graph{10: emptyNode(), 20: emptyNode(), 30: emptyNode()},
			true,
		},
		{
			"adding connection adds connection correctly",
			Graph{10: emptyNode(), 20: emptyNode(), 30: emptyNode()},
			args{10, 20, 7},
			&Graph{10: &GNode{map[int]int{20: 7}}, 20: emptyNode(), 30: emptyNode()},
			false,
		},
		{
			"adding connection overrides existing connection",
			Graph{10: &GNode{map[int]int{20: 7}}, 20: emptyNode(), 30: emptyNode()},
			args{10, 20, 1},
			&Graph{10: &GNode{map[int]int{20: 1}}, 20: emptyNode(), 30: emptyNode()},
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

func TestGraph_BreadthFirst_Errors(t *testing.T) {
	type args struct {
		valueFrom int
		valueTo   int
	}
	tests := []struct {
		name    string
		graph   Graph
		args    args
		want    []int
		wantErr bool
	}{
		{
			"valueFrom not in the graph",
			Graph{10: emptyNode(), 20: emptyNode()},
			args{5, 20},
			make([]int, 0),
			true,
		},
		{
			"valueTo not in the graph",
			Graph{10: emptyNode(), 20: emptyNode()},
			args{10, 25},
			make([]int, 0),
			true,
		},
		{
			"no path between nodes",
			Graph{10: emptyNode(), 20: emptyNode()},
			args{10, 20},
			make([]int, 0),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.graph.BreadthFirst(tt.args.valueFrom, tt.args.valueTo)
			if (err != nil) != tt.wantErr {
				t.Errorf("BreadthFirst() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BreadthFirst() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_BreadthFirst_CorrectOrder(t *testing.T) {
	type args struct {
		valueFrom int
		valueTo   int
	}
	tests := []struct {
		name  string
		graph Graph
		args  args
		want  []int
	}{
		{
			"valueTo is equal to valueFrom",
			Graph{
				10: &GNode{map[int]int{20: 1}},
				20: &GNode{map[int]int{30: 1}},
				30: emptyNode(),
			},
			args{10, 10},
			[]int{10},
		},
		{
			"3 nodes one by one",
			Graph{
				10: &GNode{map[int]int{20: 1}},
				20: &GNode{map[int]int{30: 1}},
				30: emptyNode(),
			},
			args{10, 30},
			[]int{10, 20, 30},
		},
		{
			"4 nodes in circle",
			Graph{
				10: &GNode{map[int]int{20: 1, 40: 1}},
				20: &GNode{map[int]int{30: 1}},
				30: &GNode{map[int]int{40: 1}},
				40: &GNode{map[int]int{10: 1}},
			},
			args{10, 40},
			[]int{10, 40},
		},
		{
			"6 nodes",
			Graph{
				10: &GNode{map[int]int{20: 1, 30: 1}},
				20: &GNode{map[int]int{40: 1, 50: 1, 10: 1}},
				30: &GNode{map[int]int{20: 1, 40: 1}},
				40: &GNode{map[int]int{50: 1, 20: 1}},
				50: &GNode{map[int]int{60: 1}},
				60: emptyNode(),
			},
			args{10, 60},
			[]int{10, 20, 50, 60},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.graph.BreadthFirst(tt.args.valueFrom, tt.args.valueTo)
			if err != nil {
				t.Errorf("BreadthFirst() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BreadthFirst() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_Dijkstra_Errors(t *testing.T) {
	type args struct {
		valueFrom int
		valueTo   int
	}
	tests := []struct {
		name    string
		graph   Graph
		args    args
		want    []int
		wantErr bool
	}{
		{
			"valueFrom not in the graph",
			Graph{10: emptyNode(), 20: emptyNode()},
			args{5, 20},
			make([]int, 0),
			true,
		},
		{
			"valueTo not in the graph",
			Graph{10: emptyNode(), 20: emptyNode()},
			args{10, 25},
			make([]int, 0),
			true,
		},
		{
			"no path between nodes",
			Graph{10: emptyNode(), 20: emptyNode()},
			args{10, 20},
			make([]int, 0),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.graph.Dijkstra(tt.args.valueFrom, tt.args.valueTo)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dijkstra() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dijkstra() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_Dijkstra_correctOrder(t *testing.T) {
	type args struct {
		valueFrom int
		valueTo   int
	}
	tests := []struct {
		name  string
		graph Graph
		args  args
		want  []int
	}{
		{
			"valueTo is equal to valueFrom",
			Graph{
				10: &GNode{map[int]int{20: 1}},
				20: &GNode{map[int]int{30: 2}},
				30: emptyNode(),
			},
			args{10, 10},
			[]int{10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.graph.BreadthFirst(tt.args.valueFrom, tt.args.valueTo)
			if err != nil {
				t.Errorf("BreadthFirst() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BreadthFirst() got = %v, want %v", got, tt.want)
			}
		})
	}
}
