package perm_test

import (
	"fmt"

	"github.com/bmizerany/perm"
)

func ExampleSlices() {
	slice := []int{1, 2, 3}
	p := perm.Slices(func(i, j int) bool {
		return slice[i] < slice[j]
	}, slice)
	for {
		fmt.Println(slice)
		if !p.Next() {
			break
		}
	}
	// Output:
	// [1 2 3]
	// [1 3 2]
	// [2 1 3]
	// [2 3 1]
	// [3 1 2]
	// [3 2 1]
}

func ExampleSlices_embeded() {
	order := []byte("abc")

	slice := []interface{}{
		order, // permutates with slice[1:]
		1, 2, 3,
	}

	less := func(i, j int) bool {
		return order[i] < order[j]
	}

	p := perm.Slices(less, order, slice[1:])
	for {
		fmt.Printf("%q\n", slice)
		if !p.Next() {
			break
		}
	}

	// Output:
	// ["abc" '\x01' '\x02' '\x03']
	// ["acb" '\x01' '\x03' '\x02']
	// ["bac" '\x02' '\x01' '\x03']
	// ["bca" '\x02' '\x03' '\x01']
	// ["cab" '\x03' '\x01' '\x02']
	// ["cba" '\x03' '\x02' '\x01']
}
