package _703

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		k    int
		nums []int
		want int
	}{
		{
			name: "Test Case 1",
			k:    3,
			nums: []int{4, 5, 8, 2},
			want: 4, // The 3rd largest element in [4, 5, 8, 2] is 4
		},
		{
			name: "Test Case 2",
			k:    1,
			nums: []int{4, 5, 8, 2},
			want: 8, // The largest element in [4, 5, 8, 2] is 8
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kthLargest := Constructor(tt.k, tt.nums)
			if got := kthLargest.Add(tt.nums[len(tt.nums)-1]); got != tt.want {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		h    IntHeap
		want int
	}{
		{
			name: "Test Case 1",
			h:    IntHeap{1, 2, 3},
			want: 3,
		},
		{
			name: "Test Case 2",
			h:    IntHeap{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntHeap_Less(t *testing.T) {
	tests := []struct {
		name string
		h    IntHeap
		i    int
		j    int
		want bool
	}{
		{
			name: "Test Case 1",
			h:    IntHeap{1, 2, 3},
			i:    0,
			j:    1,
			want: true,
		},
		{
			name: "Test Case 2",
			h:    IntHeap{3, 2, 1},
			i:    0,
			j:    2,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.i, tt.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    IntHeap
		want interface{}
	}{
		{
			name: "Test Case 1",
			h:    IntHeap{1, 2, 3},
			want: 3,
		},
		{
			name: "Test Case 2",
			h:    IntHeap{5},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntHeap_Push(t *testing.T) {
	tests := []struct {
		name string
		h    IntHeap
		x    int
		want IntHeap
	}{
		{
			name: "Test Case 1",
			h:    IntHeap{1, 2, 3},
			x:    4,
			want: IntHeap{1, 2, 3, 4},
		},
		{
			name: "Test Case 2",
			h:    IntHeap{},
			x:    5,
			want: IntHeap{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(tt.x)
			if !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("Push() = %v, want %v", tt.h, tt.want)
			}
		})
	}
}

func TestIntHeap_Swap(t *testing.T) {
	tests := []struct {
		name string
		h    IntHeap
		i    int
		j    int
		want IntHeap
	}{
		{
			name: "Test Case 1",
			h:    IntHeap{1, 2, 3},
			i:    0,
			j:    2,
			want: IntHeap{3, 2, 1},
		},
		{
			name: "Test Case 2",
			h:    IntHeap{5, 10, 15},
			i:    1,
			j:    2,
			want: IntHeap{5, 15, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.i, tt.j)
			if !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("Swap() = %v, want %v", tt.h, tt.want)
			}
		})
	}
}

func TestKthLargest_Add(t *testing.T) {
	tests := []struct {
		name     string
		k        int
		nums     []int
		val      int
		expected int
	}{
		{
			name:     "Test Case 1",
			k:        3,
			nums:     []int{4, 5, 8, 2},
			val:      3,
			expected: 4,
		},
		{
			name:     "Test Case 2",
			k:        1,
			nums:     []int{4, 5, 8, 2},
			val:      10,
			expected: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kthLargest := Constructor(tt.k, tt.nums)
			if got := kthLargest.Add(tt.val); got != tt.expected {
				t.Errorf("Add() = %v, want %v", got, tt.expected)
			}
		})
	}
}
