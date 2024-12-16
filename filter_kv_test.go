package iterx_test

import (
	"maps"
	"strconv"
	"testing"

	"github.com/ndx-technologies/iterx"
)

func TestFilterKV(t *testing.T) {
	tests := []struct {
		m    map[string]int
		keep func(k string, v int) bool
		e    map[string]int
	}{
		{
			m: map[string]int{"a": 1, "b": 2, "c": 3, "d": 3},
			keep: func(k string, v int) bool {
				m := map[string]bool{
					"a": true,
					"d": true,
				}
				return m[k]
			},
			e: map[string]int{"a": 1, "d": 3},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := maps.Collect(iterx.FilterKV(maps.All(tt.m), tt.keep))
			if !maps.Equal(got, tt.e) {
				t.Errorf("got %v, want %v", got, tt.e)
			}
		})
	}
}
