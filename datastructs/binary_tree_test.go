package datastructs

import (
	"reflect"
	"testing"
)

func TestMakeBSTree(t *testing.T) {
	type args struct {
		rootVal int
	}

	tests := []struct {
		name string
		args args
		want BSTree
	}{
		{
			"invalid root node",
			args{rootVal: 5},
			BSTree{RootNode: BSTNode{Value: 5, LeftChild: nil, RightChild: nil, Parent: nil}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeBSTree(tt.args.rootVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeBSTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Insert(t *testing.T) {
	testTrees := [2]BSTree{MakeBSTree(75), MakeBSTree(75)}

	tests := []struct {
		name    string
		tree    *BSTree
		inserts []int
		want    []int
	}{
		{
			"nodes are inserted correctly",
			&testTrees[0],
			[]int{25, 100, 12, 500, 50, 95},
			[]int{12, 25, 50, 75, 95, 100, 500},
		},
		{
			"nodes that are already present are not inserted",
			&testTrees[1],
			[]int{25, 25, 75},
			[]int{25, 75},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, ins := range tt.inserts {
				tt.tree.Insert(ins)
			}
			if got := tt.tree.ReturnInOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("actual in order nodes after Insert = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Delete_RootNode(t *testing.T) {
	testTrees := [4]BSTree{MakeBSTree(75), MakeBSTree(75), MakeBSTree(75), MakeBSTree(75)}

	tests := []struct {
		name  string
		tree  *BSTree
		value int
		want  []int
	}{
		{
			"root node has no children",
			&testTrees[0],
			75,
			[]int{75},
		},
		{
			"root node has left child",
			(&testTrees[1]).Insert(10).Insert(5).Insert(15),
			75,
			[]int{5, 10, 15},
		},
		{
			"root node has right child",
			(&testTrees[2]).Insert(100).Insert(95).Insert(105),
			75,
			[]int{95, 100, 105},
		},
		{
			"root node has both children",
			(&testTrees[3]).Insert(10).Insert(100).Insert(5).Insert(15).Insert(95),
			75,
			[]int{5, 10, 15, 95, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Delete(tt.value).ReturnInOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("actual in order nodes after Delete = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Delete_NonRootNode(t *testing.T) {
	testTrees := [4]BSTree{MakeBSTree(75), MakeBSTree(75), MakeBSTree(75), MakeBSTree(75)}

	tests := []struct {
		name  string
		tree  *BSTree
		value int
		want  []int
	}{
		{
			"node has no children",
			(&testTrees[0]).Insert(10),
			10,
			[]int{75},
		},
		{
			"node has left child",
			(&testTrees[1]).Insert(10).Insert(5).Insert(2).Insert(6),
			10,
			[]int{2, 5, 6, 75},
		},
		{
			"node has right child",
			(&testTrees[2]).Insert(10).Insert(15).Insert(12).Insert(20),
			10,
			[]int{12, 15, 20, 75},
		},
		{
			"node has both children",
			(&testTrees[3]).Insert(10).Insert(5).Insert(15).Insert(2).Insert(6).Insert(12).Insert(20),
			10,
			[]int{2, 5, 6, 12, 15, 20, 75},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Delete(tt.value).ReturnInOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("actual in order nodes after Delete = %v, want %v", got, tt.want)
			}
		})
	}
}
