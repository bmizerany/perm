package perm

import (
	"reflect"
	"testing"

	"github.com/bradfitz/lesser"
	"github.com/google/go-cmp/cmp"
)

type T struct {
	V string
}

func TestSlice(t *testing.T) {
	tests := []struct {
		in, want interface{}
	}{
		{
			in: []byte{1, 2, 3, 4},
			want: [][]byte{
				{1, 2, 3, 4},
				{1, 2, 4, 3},
				{1, 3, 2, 4},
				{1, 3, 4, 2},
				{1, 4, 2, 3},
				{1, 4, 3, 2},
				{2, 1, 3, 4},
				{2, 1, 4, 3},
				{2, 3, 1, 4},
				{2, 3, 4, 1},
				{2, 4, 1, 3},
				{2, 4, 3, 1},
				{3, 1, 2, 4},
				{3, 1, 4, 2},
				{3, 2, 1, 4},
				{3, 2, 4, 1},
				{3, 4, 1, 2},
				{3, 4, 2, 1},
				{4, 1, 2, 3},
				{4, 1, 3, 2},
				{4, 2, 1, 3},
				{4, 2, 3, 1},
				{4, 3, 1, 2},
				{4, 3, 2, 1},
			},
		},
		{
			in: []string{"a", "b", "c"},
			want: [][]string{
				{"a", "b", "c"},
				{"a", "c", "b"},
				{"b", "a", "c"},
				{"b", "c", "a"},
				{"c", "a", "b"},
				{"c", "b", "a"},
			},
		},
		{
			in: []rune("abc"),
			want: [][]rune{
				{'a', 'b', 'c'},
				{'a', 'c', 'b'},
				{'b', 'a', 'c'},
				{'b', 'c', 'a'},
				{'c', 'a', 'b'},
				{'c', 'b', 'a'},
			},
		},
		{
			in: []T{
				{"a"}, {"b"}, {"c"},
			},
			want: [][]T{
				{{"a"}, {"b"}, {"c"}},
				{{"a"}, {"c"}, {"b"}},
				{{"b"}, {"a"}, {"c"}},
				{{"b"}, {"c"}, {"a"}},
				{{"c"}, {"a"}, {"b"}},
				{{"c"}, {"b"}, {"a"}},
			},
		},
	}
	for _, tt := range tests {
		p := Slices(lesser.Of(tt.in), tt.in)
		t.Run("TODO", func(t *testing.T) {
			var got interface{}
			f := func(p interface{}) {
				got = appendCopy(got, p)
			}
			f(tt.in)
			for p.Next() {
				f(tt.in)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("error mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func appendCopy(s interface{}, a interface{}) interface{} {
	if s == nil {
		s = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(a)), 0, 0).Interface()
	}
	return reflect.Append(reflect.ValueOf(s), copySlice(a)).Interface()
}

func copySlice(s interface{}) reflect.Value {
	t, v := reflect.TypeOf(s), reflect.ValueOf(s)
	c := reflect.MakeSlice(t, v.Len(), v.Len())
	reflect.Copy(c, v)
	return c
}

func TestNextAllocs(t *testing.T) {
	r := []rune("0123456789abcdefghijklmnopqrstuvwxyz") // more than enough perms
	p := Slices(func(i, j int) bool {
		return r[i] < r[j]
	}, r)
	allocs := testing.AllocsPerRun(1000000, func() {
		if !p.Next() {
			t.Fatal("unexpected exhaustion of permutations")
		}
	})
	if allocs > 0 {
		t.Fatalf("allocs = %f; want 0", allocs)
	}
}

func BenchmarkNext(b *testing.B) {
	r := []rune("0123456789abcdefghijklmnopqrstuvwxyz") // more than enough perms
	p := Slices(func(i, j int) bool {
		return r[i] < r[j]
	}, r)
	for i := 0; i < b.N; i++ {
		if !p.Next() {
			b.Fatal("unexpected exhaustion of permutations")
		}
	}
}
