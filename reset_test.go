package perm

import (
	"fmt"
	"strings"
	"testing"
)

func TestReset(t *testing.T) {
	p := NewPerm(3)

	for i := 0; i < 3; i++ {
		p.Reset(3)

		var b strings.Builder
		for {
			p.Visit(func(i int) {
				fmt.Fprint(&b, i)
			})
			fmt.Fprint(&b, ":")
			if !p.Next() {
				break
			}
		}

		want := `012:021:102:120:201:210:`

		if got := b.String(); want != got {
			t.Errorf("got = %v; want %v", got, want)
		}
	}
}

func TestResetAllocs(t *testing.T) {
	const (
		n          = 3
		iterations = 10000
		perms      = 6 // == 3!
	)

	var p Perm
	var g int

	allocs := testing.AllocsPerRun(iterations, func() {
		p.Reset(n)
		for {
			p.Visit(func(i int) {
				g++
			})
			if !p.Next() {
				break
			}
		}
	})
	if allocs > 0 {
		t.Errorf("allocs = %f; want 0", allocs)
	}

	w := perms * n * (iterations + 1) // == 180018 (includes warm-up run)
	if g != w {
		t.Errorf("g = %d; want %d", g, w)
	}
}
