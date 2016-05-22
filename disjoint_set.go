package ds

// DisjointSet  represents  a  data  structure  that  stores  information  about
// connections in  a graph and  allows to determine  if two arbitrary  nodes are
// connected in O(log n).
type DisjointSet struct {
	id   []int
	size []int
}

// NewDisjointSet returns a new DisjointSet initialized to n elements.
func NewDisjointSet(n uint) *DisjointSet {
	ds := &DisjointSet{
		id:   make([]int, n),
		size: make([]int, n),
	}
	for i := 0; i < int(n); i++ {
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

// Connected determines if nodes u and v are connected in some way.
func (t *DisjointSet) Connected(u, v int) bool {
	return t.root(u) == t.root(v)
}

// Connect connects nodes u and v.
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
