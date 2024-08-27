package _1514

import (
	"container/heap"
)

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	// 2 dim array with map
	graph := make([]map[int]float64, n)
	for i := range graph {
		graph[i] = make(map[int]float64)
	}
	//maxHeap as priority queue
	maxHeap := make(PriorityQueue, 0)
	//init graph
	maxHeap.Push(&Item{value: start_node, priority: 1})

	seen := make([]bool, n)

	for i := range edges {
		u := edges[i][0]
		v := edges[i][1]
		prob := succProb[i]
		//graph = append(graph, map[int]float64{v: prob})
		//graph = append(graph, map[int]float64{u: prob})
		graph[u][v] = prob
		graph[v][u] = prob

	}

	for maxHeap.Len() > 0 {
		item := heap.Pop(&maxHeap).(*Item)
		prob := item.priority
		u := item.value

		if u == end_node {
			return prob
		}

		if seen[u] {
			continue
		}

		seen[u] = true

		for nextNode, edgeProb := range graph[u] {
			if seen[nextNode] {
				continue
			}
			heap.Push(&maxHeap, &Item{value: nextNode, priority: prob * edgeProb})
		}
	}

	return 0.0
}
