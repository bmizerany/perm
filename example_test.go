package perm_test

import (
	"fmt"

	"blake.io/perm"
)

func ExamplePerm() {
	animals := []string{"bird", "cat", "fish"}
	p := perm.NewPerm(len(animals))

	for {
		fmt.Print("animals:")
		p.Visit(func(i int) {
			fmt.Print(" ", animals[i])
		})
		fmt.Println()
		if !p.Next() {
			break
		}
	}

	// Output:
	// animals: bird cat fish
	// animals: bird fish cat
	// animals: cat bird fish
	// animals: cat fish bird
	// animals: fish bird cat
	// animals: fish cat bird
}

func ExampleComb() {
	animals := []string{"bird", "cat", "fish"}
	c := perm.NewComb(len(animals))
	for {
		fmt.Print("animals:")
		c.Visit(func(i int) {
			fmt.Print(" ", animals[i])
		})
		fmt.Println()

		if !c.Next() {
			break
		}
	}

	// Output:
	// animals:
	// animals: bird
	// animals: bird cat
	// animals: cat
	// animals: cat fish
	// animals: bird cat fish
	// animals: bird fish
	// animals: fish
}

// The empty subset is always visited first, so if
// you need to visit all nonempty subsets, check c.Next() at the
// beginning of the loop instead of at the end.
func ExampleComb_nonempty() {
	animals := []string{"bird", "cat", "fish"}
	c := perm.NewComb(len(animals))
	for c.Next() {
		fmt.Print("animals:")
		c.Visit(func(i int) {
			fmt.Print(" ", animals[i])
		})
		fmt.Println()
	}

	// Output:
	// animals: bird
	// animals: bird cat
	// animals: cat
	// animals: cat fish
	// animals: bird cat fish
	// animals: bird fish
	// animals: fish
}

func ExampleCombPerm() {
	animals := []string{"bird", "cat", "fish"}
	b := perm.NewCombPerm(len(animals))

	for {
		fmt.Print("animals:")
		b.Visit(func(i int) {
			fmt.Print(" ", animals[i])
		})
		fmt.Println()
		if !b.Next() {
			break
		}
	}

	// Output:
	// animals:
	// animals: bird
	// animals: bird cat
	// animals: cat bird
	// animals: cat
	// animals: cat fish
	// animals: fish cat
	// animals: bird cat fish
	// animals: bird fish cat
	// animals: cat bird fish
	// animals: cat fish bird
	// animals: fish bird cat
	// animals: fish cat bird
	// animals: bird fish
	// animals: fish bird
	// animals: fish
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
