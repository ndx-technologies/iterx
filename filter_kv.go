package iterx

import "iter"

func FilterKV[K, V comparable](i iter.Seq2[K, V], keep func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range i {
			if keep(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
