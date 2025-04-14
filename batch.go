package iterx

import "iter"

// Batch iterator into groups of n elements.
// Returned slices is reused between iterators. If caller needs a copy, caller should copy himself.
func Batch[T any](it iter.Seq[T], n int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		batch := make([]T, 0, n)

		for v := range it {
			batch = append(batch, v)

			if len(batch) == n {
				if !yield(batch) {
					return
				}

				batch = batch[:0]
			}
		}
		if len(batch) > 0 {
			yield(batch)
		}
	}
}
