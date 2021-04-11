package datastructs

import (
	"reflect"
	"testing"
)

func TestMakeTree(t *testing.T) {
	type args struct {
		rootVal int
	}

	tests := []struct {
		name string
		args args
		want Tree
	}{
		{
			"invalid root node",
			args{rootVal: 5},
			Tree{RootNode: Node{Value: 5, LeftChild: nil, RightChild: nil, Parent: nil}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeTree(tt.args.rootVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Insert(t *testing.T) {
	t.Run("nodes are inserted correctly", func(t *testing.T) {
		tree := MakeTree(75)
		tree.Insert(25)
		tree.Insert(100)
		tree.Insert(12)
		tree.Insert(500)
		tree.Insert(50)
		tree.Insert(95)

		want := []int{12, 25, 50, 75, 95, 100, 500}

		if got := tree.ReturnInOrder(); !reflect.DeepEqual(got, want) {
			t.Errorf("actual in order nodes after inserts = %v, want %v", got, want)
		}
	})

	t.Run("nodes that are already present are not inserted", func(t *testing.T) {
		tree := MakeTree(75)
		tree.Insert(25)
		tree.Insert(25)
		tree.Insert(75)

		want := []int{25, 75}

		if got := tree.ReturnInOrder(); !reflect.DeepEqual(got, want) {
			t.Errorf("actual in order nodes after inserts = %v, want %v", got, want)
		}
	})
}
