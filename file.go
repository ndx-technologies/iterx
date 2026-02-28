package iterx

import (
	"io"
	"iter"
	"log/slog"
	"os"
)

func FromFile[T any](path string, parse func(io.Reader) iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		f, err := os.Open(path)
		if err != nil {
			slog.Error("failed to open file", "path", path, "error", err)
			return
		}
		defer f.Close()

		for v := range parse(f) {
			if !yield(v) {
				return
			}
		}
	}
}
