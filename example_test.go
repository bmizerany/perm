package perm_test

import (
	"fmt"

	"github.com/bmizerany/perm"
)

func ExamplePerm() {
	slice := []int{1, 2, 3}

	p := perm.NewPerm(len(slice))

	for {
		fmt.Print("order:")
		p.Visit(func(i int) {
			fmt.Print(" ", slice[i])
		})
		fmt.Println()
		if !p.Next() {
			break
		}
	}

	// Output:
	// order: 1 2 3
	// order: 1 3 2
	// order: 2 1 3
	// order: 2 3 1
	// order: 3 1 2
	// order: 3 2 1
}

func ExampleComb() {
	s := []int{1, 2, 3}
	c := perm.NewComb(len(s))
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
	c := perm.NewComb(len(s))
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
	c := perm.NewComb(len(s))

	var g []int
	for c.Next() {
		g := g[:0]
		c.Visit(func(i int) {
			g = append(g, s[i])
		})

		p := perm.NewPerm(len(g))
		for {
			fmt.Print("order:")
			p.Visit(func(i int) {
				fmt.Print(" ", g[i])
			})
			fmt.Println()
			if !p.Next() {
				break
			}
		}
	}

	// Output:
	// order: 1
	// order: 1 2
	// order: 2 1
	// order: 2
	// order: 2 3
	// order: 3 2
	// order: 1 2 3
	// order: 1 3 2
	// order: 2 1 3
	// order: 2 3 1
	// order: 3 1 2
	// order: 3 2 1
	// order: 1 3
	// order: 3 1
	// order: 3
}

// TODO: include example permutating multiple slices in step.
