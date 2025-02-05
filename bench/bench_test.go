package bench

import "testing"

func Benchmark(b *testing.B) {
	for b.Loop() {
		sum(1, 2, 3, 4, 5)
	}
}
