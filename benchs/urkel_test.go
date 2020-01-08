package benchs

import (
	"testing"
)

func BenchmarkUrkelBlock100(b *testing.B) {
	runBlockBenchmark(b, newUrkel, 100, path("mem-urkel-100.out"), path("time-urkel-100.out"))
}

func BenchmarkUrkelBlock1000(b *testing.B) {
	runBlockBenchmark(b, newUrkel, 1000, path("mem-urkel-1000.out"), path("time-urkel-1000.out"))
}

func BenchmarkUrkelBlock10000(b *testing.B) {
	runBlockBenchmark(b, newUrkel, 10000, path("mem-urkel-10000.out"), path("time-urkel-10000.out"))
}

func BenchmarkUrkelBlock40000(b *testing.B) {
	runBlockBenchmark(b, newUrkel, 40000, path("mem-urkel-40000.out"), path("time-urkel-40000.out"))
}
