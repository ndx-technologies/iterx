package iterx

import "iter"

func MergeOrdered[T any](cmp func(T, T) bool, iters ...iter.Seq[T]) iter.Seq[T] {
	next := make([]func() (T, bool), 0, len(iters))
	stop := make([]func(), 0, len(iters))

	for _, i := range iters {
		if i == nil {
			continue
		}
		n, s := iter.Pull(i)
		next = append(next, n)
		stop = append(stop, s)
	}

	return func(yield func(T) bool) {
		for _, stop := range stop {
			defer stop()
		}

		vs := make([]T, len(iters))
		ok := make([]bool, len(iters))

		for i, next := range next {
			vs[i], ok[i] = next()
		}

		for {
			var (
				min    T
				minIdx int
				found  bool
			)

			for i, v := range vs {
				if !ok[i] {
					continue
				}

				if !found || cmp(v, min) {
					min = v
					minIdx = i
					found = true
				}
			}

			if !found {
				return
			}

			if !yield(min) {
				return
			}

			vs[minIdx], ok[minIdx] = next[minIdx]()
		}
	}
}
