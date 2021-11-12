package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"perm"
	"sort"
	"strings"
)

var (
	flagComb = flag.Bool("c", false, "include combinations")
	flagPerm = flag.Bool("p", false, "include permutations")
	flagSep  = flag.String("s", " ", "separation character. default is space")
)

func main() {
	flag.Parse()

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	items := strings.Split(string(input), "\n")
	items = trimSpaceAll(items)
	sort.Strings(items)

	var p interface {
		Next() bool
		Visit(func(i int))
	}

	switch {
	case *flagPerm && *flagComb:
		p = perm.NewCombPerm(len(items))
	case *flagComb:
		p = perm.NewComb(len(items))
	default:
		p = perm.NewPerm(len(items))
	}

	for {
		n := 0
		p.Visit(func(i int) {
			if n > 0 {
				io.WriteString(os.Stdout, *flagSep)
			}
			n++
			fmt.Fprintf(os.Stdout, "%s", items[i])
		})

		if n > 0 {
			io.WriteString(os.Stdout, "\n")
		}

		if !p.Next() {
			break
		}
	}
}

func trimSpaceAll(ss []string) []string {
	var out []string
	for _, s := range ss {
		s = strings.TrimSpace(s)
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
