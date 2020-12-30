package perm

type Both struct {
	c *Comb
	p *Perm
	v []int
}

func NewBoth(n int) *Both {
	return &Both{c: NewComb(n), p: NewPerm(0)}
}

func (b *Both) Next() bool {
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
	b.p = NewPerm(len(b.v))
	return true
}

func (b *Both) Visit(f func(i int)) {
	b.p.Visit(func(i int) {
		f(b.v[i])
	})
}
