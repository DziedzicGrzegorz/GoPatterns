package _703

import (
	"container/heap"
)

// KthLargest struct
type KthLargest struct {
	minHeap *IntHeap
	k       int
}

// IntHeap is a min-heap of integers
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Constructor initializes the KthLargest struct with k and nums
func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)

	kthLargest := KthLargest{
		minHeap: h,
		k:       k,
	}

	for _, num := range nums {
		kthLargest.Add(num)
	}

	return kthLargest
}

// Add adds a new element to the stream and returns the kth largest element
func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)
	if this.minHeap.Len() > this.k {
		heap.Pop(this.minHeap)
	}
	return (*this.minHeap)[0]
}
