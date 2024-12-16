package iterx

import "iter"

func MergeKV[K comparable, V any](is ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	if len(is) == 0 {
		return func(func(K, V) bool) {}
	}
	return func(yield func(K, V) bool) {
		for _, i := range is {
			if i == nil {
				continue
			}
			for k, v := range i {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
