package perm

import (
	"fmt"
	"strings"
	"testing"
)

func TestReset(t *testing.T) {
	type I interface {
		Next() bool
		Visit(func(i int))
		Reset(int)
	}

	tests := []struct {
		x I
		w string
	}{
		{NewPerm(0), "012.021.102.120.201.210.01.10.0.."},
		{NewComb(0), ".0.01.1.12.012.02.2..0.01.1..0.."},
		{NewCombPerm(0), ".0.01.10.1.12.21.012.021.102.120.201.210.02.20.2..0.01.10.1..0.."},
	}

	for _, tt := range tests {
		var b strings.Builder
		for i := 3; i >= 0; i-- {
			p := tt.x

			p.Reset(i)

			for {
				p.Visit(func(i int) {
					fmt.Fprint(&b, i)
				})
				fmt.Fprint(&b, ".")
				if !p.Next() {
					break
				}
			}

		}
		if got := b.String(); tt.w != got {
			t.Errorf("%T: got = %v; want %v", tt.x, got, tt.w)
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
