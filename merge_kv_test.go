package iterx_test

import (
	"iter"
	"maps"
	"strconv"
	"testing"

	"github.com/ndx-technologies/iterx"
)

func TestMergeKV(t *testing.T) {
	tests := []struct {
		ms []map[string]int
		e  map[string]int
	}{
		{
			ms: []map[string]int{
				{"a": 1, "b": 2, "c": 3, "d": 3},
				{"a": 2, "e": 4},
				{"g": 13},
				{},
				nil,
			},
			e: map[string]int{"a": 2, "b": 2, "c": 3, "d": 3, "e": 4, "g": 13},
		},
		{
			ms: []map[string]int{nil, nil, nil, {}},
			e:  nil,
		},
		{
			ms: nil,
			e:  nil,
		},
	}
	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var iters []iter.Seq2[string, int]
			for _, m := range tc.ms {
				iters = append(iters, maps.All(m))
			}

			merged := make(map[string]int)
			for k, v := range iterx.MergeKV(iters...) {
				merged[k] = v
			}

			if !maps.Equal(merged, tc.e) {
				t.Errorf("got %v, want %v", merged, tc.e)
			}
		})
	}
}
