// Package perm permutates data.
package perm

import (
	"reflect"
	"sort"
)

// Slices returns a Perm given the provided slices and less function. Each
// slice is mutated in step with the others according less.
//
// The function assumes all slices are sorted.
//
// The function panics if the provided interfaces are not slices, or will panic
// on Next if all slices in slice are not of equal length.
func Slices(less func(i, j int) bool, slice ...interface{}) Perm {
	if len(slice) == 0 {
		return Perm{}
	}

	swaps := make([]func(i, j int), len(slice))
	for i := range swaps {
		swaps[i] = reflect.Swapper(slice[i])
	}

	return Perm{
		size: reflect.ValueOf(slice[0]).Len(),
		swap: func(i, j int) {
			for _, s := range swaps {
				s(i, j)
			}
		},
		less: less,
	}
}

// Perm permutates slices.
type Perm struct {
	size int
	swap func(i, j int)
	less func(i, j int) bool
}

// New returns a new Perm that permutates data.
//
// The function assumes data is sorted.
func New(data sort.Interface) Perm {
	return Perm{
		size: data.Len(),
		swap: data.Swap,
		less: data.Less,
	}
}

// Next mutates the underlying slice to the next permutation.  See Knuth's
// Algorithm L.
//
// It does no allocate.
func (p Perm) Next() bool {
	n := p.size - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !p.less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !p.less(j, l) {
		l--
	}
	p.swap(j, l)
	for k, l := j+1, n; k < l; {
		p.swap(k, l)
		k++
		l--
	}
	return true
}
