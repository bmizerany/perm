package perm_test

import (
	"fmt"

	"github.com/bmizerany/perm"
)

func ExamplePerm() {
	p := perm.NewPerm(3)

	for {
		fmt.Print("order:")
		p.Visit(func(i int) {
			fmt.Print(" ", i)
		})
		fmt.Println()
		if !p.Next() {
			break
		}
	}

	// Output:
	// order: 0 1 2
	// order: 0 2 1
	// order: 1 0 2
	// order: 1 2 0
	// order: 2 0 1
	// order: 2 1 0
}

func ExampleComb() {
	c := perm.NewComb(3)
	for {
		fmt.Print("subset:")
		c.Visit(func(i int) {
			fmt.Print(" ", i)
		})
		fmt.Println()

		if !c.Next() {
			break
		}
	}

	// Output:
	// subset:
	// subset: 0
	// subset: 0 1
	// subset: 1
	// subset: 1 2
	// subset: 0 1 2
	// subset: 0 2
	// subset: 2
}

// The empty subset is always visited first, so if
// you need to visit all nonempty subsets, check c.Next() at the
// beginning of the loop instead of at the end.
func ExampleComb_nonempty() {
	c := perm.NewComb(3)
	for c.Next() {
		fmt.Print("subset:")
		c.Visit(func(i int) {
			fmt.Print(" ", i)
		})
		fmt.Println()
	}

	// Output:
	// subset: 0
	// subset: 0 1
	// subset: 1
	// subset: 1 2
	// subset: 0 1 2
	// subset: 0 2
	// subset: 2
}

func ExampleBoth() {
	b := perm.NewBoth(3)

	for {
		fmt.Print("order:")
		b.Visit(func(i int) {
			fmt.Print(" ", i)
		})
		fmt.Println()
		if !b.Next() {
			break
		}
	}

	// Output:
	// order:
	// order: 0
	// order: 0 1
	// order: 1 0
	// order: 1
	// order: 1 2
	// order: 2 1
	// order: 0 1 2
	// order: 0 2 1
	// order: 1 0 2
	// order: 1 2 0
	// order: 2 0 1
	// order: 2 1 0
	// order: 0 2
	// order: 2 0
	// order: 2
}

func ExamplePerm_multi() {
	p := perm.NewPerm(3)

	a := []byte("abc")
	b := []string{"foo", "bar", "baz"}

	for {
		p.Visit(func(i int) {
			fmt.Printf("%c", a[i])
		})
		fmt.Print(":")
		p.Visit(func(i int) {
			fmt.Print(" ", b[i])
		})
		fmt.Println()
		if !p.Next() {
			break
		}
	}

	// Output:
	// abc: foo bar baz
	// acb: foo baz bar
	// bac: bar foo baz
	// bca: bar baz foo
	// cab: baz foo bar
	// cba: baz bar foo
}
