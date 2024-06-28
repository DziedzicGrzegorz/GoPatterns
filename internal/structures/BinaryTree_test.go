package structures

import (
	"reflect"
	"testing"
)

func TestNode_InOrderTraversal(t *testing.T) {
	tests := []struct {
		name   string
		fields *Node
		want   []int
	}{
		{
			name:   "Single node",
			fields: &Node{Value: 10},
			want:   []int{10},
		},
		{
			name:   "Left child only",
			fields: &Node{Value: 10, Left: &Node{Value: 5}},
			want:   []int{5, 10},
		},
		{
			name:   "Right child only",
			fields: &Node{Value: 10, Right: &Node{Value: 15}},
			want:   []int{10, 15},
		},
		{
			name:   "Both children",
			fields: &Node{Value: 10, Left: &Node{Value: 5}, Right: &Node{Value: 15}},
			want:   []int{5, 10, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			tt.fields.InOrderTraversal(func(val int) {
				got = append(got, val)
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InOrderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Insert(t *testing.T) {
	tests := []struct {
		name   string
		fields *Node
		args   int
		want   *Node
	}{
		{
			name:   "Insert to empty tree",
			fields: &Node{Value: 10},
			args:   5,
			want:   &Node{Value: 10, Left: &Node{Value: 5}},
		},
		{
			name:   "Insert to left child",
			fields: &Node{Value: 10, Left: &Node{Value: 5}},
			args:   3,
			want:   &Node{Value: 10, Left: &Node{Value: 5, Left: &Node{Value: 3}}},
		},
		{
			name:   "Insert to right child",
			fields: &Node{Value: 10, Right: &Node{Value: 15}},
			args:   20,
			want:   &Node{Value: 10, Right: &Node{Value: 15, Right: &Node{Value: 20}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tt.fields
			n.Insert(tt.args)
			if !reflect.DeepEqual(n, tt.want) {
				t.Errorf("Insert() = %v, want %v", n, tt.want)
			}
		})
	}
}

func TestNode_Search(t *testing.T) {
	tests := []struct {
		name   string
		fields *Node
		args   int
		want   bool
	}{
		{
			name:   "Search in empty tree",
			fields: &Node{Value: 10},
			args:   5,
			want:   false,
		},
		{
			name:   "Search existing value",
			fields: &Node{Value: 10, Left: &Node{Value: 5}},
			args:   5,
			want:   true,
		},
		{
			name:   "Search non-existing value",
			fields: &Node{Value: 10, Left: &Node{Value: 5}},
			args:   15,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tt.fields
			if got := n.Search(tt.args); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
