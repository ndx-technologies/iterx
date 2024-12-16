package iterx_test

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/ndx-technologies/iterx"
)

func TestMergeOrdered(t *testing.T) {
	type S struct {
		c string
		v int
		t int
	}

	cmp := func(a, b S) bool { return a.t < b.t }

	tests := []struct {
		vs     [][]S
		merged []S
	}{
		{
			vs: [][]S{
				{{c: "r", v: 1, t: 1}, {c: "r", v: 3, t: 2}, {c: "r", v: 5, t: 5}},
				{{c: "g", v: 2, t: 3}},
				{{c: "b", v: 0, t: 4}, {c: "g", v: 7, t: 7}, {c: "g", v: 8, t: 9}},
			},
			merged: []S{
				{c: "r", v: 1, t: 1},
				{c: "r", v: 3, t: 2},
				{c: "g", v: 2, t: 3},
				{c: "b", v: 0, t: 4},
				{c: "r", v: 5, t: 5},
				{c: "g", v: 7, t: 7},
				{c: "g", v: 8, t: 9},
			},
		},
		{
			vs: [][]S{
				{{c: "r", v: 1, t: 1}, {c: "r", v: 3, t: 2}, {c: "r", v: 5, t: 5}},
				nil,
				nil,
			},
			merged: []S{
				{c: "r", v: 1, t: 1},
				{c: "r", v: 3, t: 2},
				{c: "r", v: 5, t: 5},
			},
		},
		{
			vs: [][]S{
				nil,
				nil,
				nil,
			},
			merged: nil,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			iters := make([]iter.Seq[S], len(test.vs))
			for i, vs := range test.vs {
				iters[i] = slices.Values(vs)
			}

			var merged []S
			for q := range iterx.MergeOrdered(cmp, iters...) {
				merged = append(merged, q)
			}

			total := 0
			for _, vs := range test.vs {
				total += len(vs)
			}
			if !slices.Equal(merged, test.merged) {
				t.Errorf("got %v, want %v", merged, test.merged)
			}
		})
	}
}
