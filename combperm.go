package perm

// CombPerm generates all permutations of all subsets of sequences of a
// given length.
//
// The first subset is always the empty set. To skip the empty set, call
// Next before the first call to Visit.
type CombPerm struct {
	c Comb
	p Perm
	v []int
}

// NewCombPerm returns a new CombPerm for length n.
func NewCombPerm(n int) *CombPerm {
	var c CombPerm
	c.Reset(n)
	return &c
}

// Next advances b to the next permutation.
func (b *CombPerm) Next() bool {
	if b.p.Next() {
		return true
	}

	if !b.c.Next() {
		return false
	}

	b.v = b.v[:0]
	b.c.Visit(func(i int) {
		b.v = append(b.v, i)
	})
	b.p.Reset(len(b.v))
	return true
}

// Visit calls f for each index in the current permutation.
func (b *CombPerm) Visit(f func(i int)) {
	b.p.Visit(func(i int) {
		f(b.v[i])
	})
}

func (b *CombPerm) Reset(n int) {
	b.c.Reset(n)
	b.p.Reset(0)
}
