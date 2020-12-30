package perm

// Knuth:
//
// Algorithm L (Loopless Gray binary generation)
//
// L1. [Initialize.] Set a_j <- 0 and f_j <- j for 0 <= j < n;
//     also set f_n <- n. (A loopless algorithm is allowed to
//     have loops in its initializaiton step, as long as the
//     initial setup is reasonably efficient; after all, every
//     program needs to be loaded and launched.)
// L2. [Visit.] Visit the n-tuple (a_(n-1), ..., a_1, a_0).
// L3. [Choose j.] Set j <- f_0, f_0 <- 0. (If this is the kth
//     time we are performing the present step, j is now equal
//     to ρ(k).) Terminate if j == n; otherwise set
//     f_j <- f_(j+1) and f_(j+1) <- j + 1.
// L4. [Complement coordinate j.] Set a_j <- 1 - a_j and return
//     to L2.

type Comb struct {
	j int
	f []int // "focus pointers"

	a []bool
}

func NewComb(n int) *Comb {
	var c Comb
	c.Reset(n)
	return &c
}

func (c *Comb) Next() bool {
	c.j = c.f[0]
	c.f[0] = 0
	if c.j == len(c.a) {
		return false
	}
	c.f[c.j] = c.f[c.j+1]
	c.f[c.j+1] = c.j + 1
	c.a[c.j] = !c.a[c.j]
	return true
}

func (c *Comb) Visit(f func(i int)) {
	for i, include := range c.a {
		if include {
			f(i)
		}
	}
}

func (c *Comb) Reset(n int) {
	if n+1 > cap(c.f) {
		c.f = make([]int, n+1)
		c.a = make([]bool, n)
	} else {
		c.f = c.f[:n+1]
		c.a = c.a[:n]
		for i := range c.a {
			c.a[i] = false
		}
	}
	c.j = 0
	for i := range c.f {
		c.f[i] = i
	}
}
