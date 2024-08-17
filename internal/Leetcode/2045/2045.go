package _2045

import (
	"math"
)

func secondMinimum(n int, edges [][]int, time int, change int) int {
	graph := make(map[int][]int)
	for _, edg := range edges {
		graph[edg[0]] = append(graph[edg[0]], edg[1])
		graph[edg[1]] = append(graph[edg[1]], edg[0])
	}
	queue := []Pair{Pair{cost: 0, node: 1}}
	dis1 := make([]int, n+1)
	dis2 := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dis1[i] = math.MaxInt
		dis2[i] = math.MaxInt
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, nei := range graph[node.node] {

			new_cost := time + node.cost
			// red signal
			if (node.cost/change)%2 == 1 {
				new_cost += change - (node.cost % change)
			}
			// update two distances
			if dis1[nei] > new_cost {
				dis2[nei] = dis1[nei]
				dis1[nei] = new_cost
				queue = append(queue, Pair{cost: new_cost, node: nei})
			} else if dis1[nei] < new_cost && dis2[nei] > new_cost {
				dis2[nei] = new_cost
				queue = append(queue, Pair{cost: new_cost, node: nei})
			}
		}
	}
	return dis2[n]
}

type Pair struct {
	cost int
	node int
}
