package perm

import (
	"fmt"
	"sort"
)

func ExampleComb() {
	s := []int{1, 2, 3}
	c := NewComb(len(s))
	for {
		fmt.Print("subset:")
		c.Visit(func(i int) {
			fmt.Print(" ", s[i])
		})
		fmt.Println()

		if !c.Next() {
			break
		}
	}

	// Output:
	// subset:
	// subset: 1
	// subset: 1 2
	// subset: 2
	// subset: 2 3
	// subset: 1 2 3
	// subset: 1 3
	// subset: 3
}

// The empty subset is always visited first, so if
// you need to visit all nonempty subsets, check c.Next() at the
// beginning of the loop instead of at the end.
func ExampleComb_nonempty() {
	s := []int{1, 2, 3}
	c := NewComb(len(s))
	for c.Next() {
		fmt.Print("subset:")
		c.Visit(func(i int) {
			fmt.Print(" ", s[i])
		})
		fmt.Println()
	}

	// Output:
	// subset: 1
	// subset: 1 2
	// subset: 2
	// subset: 2 3
	// subset: 1 2 3
	// subset: 1 3
	// subset: 3
}

func ExampleComb_withPerm() {
	s := []int{1, 2, 3}
	c := NewComb(len(s))

	var g sort.IntSlice
	for c.Next() {
		g := g[:0]
		c.Visit(func(i int) {
			g = append(g, s[i])
		})

		p := New(g)
		for {
			fmt.Println(g)
			if !p.Next() {
				break
			}
		}
	}

	// Output:
	// [1]
	// [1 2]
	// [2 1]
	// [2]
	// [2 3]
	// [3 2]
	// [1 2 3]
	// [1 3 2]
	// [2 1 3]
	// [2 3 1]
	// [3 1 2]
	// [3 2 1]
	// [1 3]
	// [3 1]
	// [3]
}
