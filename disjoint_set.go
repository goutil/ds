package ds

type DisjointSet struct {
	id   []int
	size []int
}

func NewDisjointSet(nodes uint) *DisjointSet {
	ds := &DisjointSet{
		id:   make([]int, nodes),
		size: make([]int, nodes),
	}
	for i := 0; i < int(nodes); i++ {
		ds.id[i] = i
		ds.size[i] = 1
	}
	return ds
}

func (t *DisjointSet) root(u int) int {
	for u != t.id[u] {
		u = t.id[u]
	}
	return u
}

func (t *DisjointSet) Connected(u, v int) bool {
	return t.root(u) == t.root(v)
}

func (t *DisjointSet) Connect(u, v int) {
	u = t.root(u)
	v = t.root(v)
	if t.size[u] < t.size[v] {
		t.id[u] = v
		t.size[v] += t.size[u]
	} else {
		t.id[v] = u
		t.size[u] += t.size[v]
	}
}
