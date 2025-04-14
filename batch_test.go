package iterx_test

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/ndx-technologies/iterx"
)

func collectBatches(it iter.Seq[[]int]) [][]int {
	var result [][]int
	for batch := range it {
		v := make([]int, len(batch))
		copy(v, batch)
		result = append(result, v)
	}
	return result
}

func TestBatch(t *testing.T) {
	tests := []struct {
		nums []int
		n    int
		e    [][]int
	}{
		{
			nums: []int{1, 2, 3, 4, 5},
			n:    2,
			e:    [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			nums: []int{1, 2},
			n:    3,
			e:    [][]int{{1, 2}},
		},
		{
			nums: []int{},
			n:    2,
			e:    [][]int{},
		},
		{
			nums: []int{1, 2, 3},
			n:    1,
			e:    [][]int{{1}, {2}, {3}},
		},
		{
			nums: []int{1, 2, 3},
			n:    5,
			e:    [][]int{{1, 2, 3}},
		},
	}
	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			batches := collectBatches(iterx.Batch(slices.Values(tc.nums), tc.n))

			if !slices.EqualFunc(batches, tc.e, slices.Equal) {
				t.Error(batches, tc.e)
			}
		})
	}
}
