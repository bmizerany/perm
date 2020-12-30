package perm

import (
	"fmt"
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
