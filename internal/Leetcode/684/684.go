package _484

type UnionFind struct {
	parent,
	rank []int
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return false
	}
	if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else if uf.rank[rootY] > uf.rank[rootX] {
		uf.parent[rootX] = rootY
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	return true
}
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rang := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
		rank:   rang,
	}
}

func findRedundantConnection(edges [][]int) []int {
	unionFind := NewUnionFind(len(edges) + 1)

	for _, edge := range edges {
		if !unionFind.Union(edge[0], edge[1]) {
			return []int{edge[0], edge[1]}
		}
	}
	return nil
}
