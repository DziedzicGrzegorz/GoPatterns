package _1334

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	adj := map[int][]int{}
	for _, edge := range edges {
		v1, v2, dist := edge[0], edge[1], edge[2]
		adj[v1] = append(adj[v1], v2, dist)
		adj[v2] = append(adj[v2], v1, dist)

	}
	res, min_count := -1, n
	for src := range n {
		count := dijkstra(n, adj, src, distanceThreshold)
		if count <= min_count {
			res, min_count = src, count
		}
	}
	return res

}
func dijkstra(n int, adj map[int][]int, src int, distanceThreshold int) int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = 1 << 31
	}
	dist[src] = 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		u := -1
		for j, v := range visited {
			if !v && (u == -1 || dist[j] < dist[u]) {
				u = j
			}
		}
		visited[u] = true
		for j := 0; j < len(adj[u]); j += 2 {
			v, w := adj[u][j], adj[u][j+1]
			if dist[v] > dist[u]+w {
				dist[v] = dist[u] + w
			}
		}
	}
	count := 0
	for _, d := range dist {
		if d <= distanceThreshold {
			count++
		}
	}
	return count
}
