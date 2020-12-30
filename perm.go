// Package perm permutes data.
package perm

// Perm generates all permutations of a given length.
type Perm struct {
	size int
	cur  []int
}

// New returns a new Perm that permutes data.
//
// The function assumes data is sorted.
func NewPerm(n int) Perm {
	p := Perm{
		size: n,
		cur:  make([]int, n),
	}
	for i := range p.cur {
		p.cur[i] = i
	}
	return p
}

func (p *Perm) Visit(f func(i int)) {
	for _, i := range p.cur {
		f(i)
	}
}

// Next advances p to the next permutation.
func (p Perm) Next() bool {
	n := p.size - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; p.cur[j] >= p.cur[j+1]; j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for p.cur[j] >= p.cur[l] {
		l--
	}
	p.cur[j], p.cur[l] = p.cur[l], p.cur[j]
	for k, l := j+1, n; k < l; {
		p.cur[k], p.cur[l] = p.cur[l], p.cur[k]
		k++
		l--
	}
	return true
}
