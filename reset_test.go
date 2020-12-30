package perm

import (
	"fmt"
	"strings"
	"testing"
)

type I interface {
	Next() bool
	Visit(func(i int))
	Reset(int)
}

func TestReset(t *testing.T) {
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

func BenchmarkResetAllocs(b *testing.B) {
	p := NewCombPerm(0)
	for i := 0; i < b.N; i++ {
		p.Reset(3)
		for {
			p.Visit(func(i int) {
				// nop
			})
			if !p.Next() {
				break
			}
		}
	}
}

func TestResetAllocs(t *testing.T) {
	const (
		n          = 3
		iterations = 10000
	)

	tests := []I{
		NewPerm(0),
		NewComb(0),
		NewCombPerm(0),
	}

	for _, p := range tests {
		allocs := int(testing.AllocsPerRun(iterations, func() {
			p.Reset(n)
			for {
				p.Visit(func(i int) {
					// nop
				})
				if !p.Next() {
					break
				}
			}
		}))
		if allocs > 0 {
			t.Errorf("%T: allocs = %d; want 0", p, allocs)
		}
	}

}
