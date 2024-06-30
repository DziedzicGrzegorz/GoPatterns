package _1579

type UnionFind struct {
	parent,
	rank []int
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX == rootY {
		return false //są już w zbiorze
	}
	if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	return true
}
func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
		rank:   rank,
	}
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	ufCommon := NewUnionFind(n + 1)
	ufFisrt := NewUnionFind(n + 1)
	ufSecond := NewUnionFind(n + 1)
	edgeToRemove := 0

	//3tego typu

	for _, edge := range edges {
		if edge[0] == 3 {
			//edge already exist so we can cut it
			if !ufCommon.Union(edge[1], edge[2]) {
				edgeToRemove++
			} else {
				ufFisrt.Union(edge[1], edge[2])
				ufSecond.Union(edge[1], edge[2])
			}
		}
	}
	for _, edge := range edges {
		if edge[0] == 1 {
			if !ufFisrt.Union(edge[1], edge[2]) {
				edgeToRemove++
			}
		} else if edge[0] == 2 {
			if !ufSecond.Union(edge[1], edge[2]) {
				edgeToRemove++
			}
		}
	}
	rootFirst := ufFisrt.find(1)
	rootSecond := ufSecond.find(1)

	for i := 2; i <= n; i++ {
		if ufFisrt.find(i) != rootFirst || ufSecond.find(i) != rootSecond {
			return -1
		}
	}
	return edgeToRemove
}
