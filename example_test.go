package perm_test

import (
	"fmt"

	"github.com/bmizerany/perm"
)

func ExamplePerm_visit() {
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

// TODO: include example permutating multiple slices in step.
