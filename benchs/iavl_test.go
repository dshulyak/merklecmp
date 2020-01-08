package benchs

import (
	"testing"
)

func BenchmarkIavlTreeBlock100(b *testing.B) {
	runBlockBenchmark(b, newIavlTree, 100, path("mem-iavlTree-100.out"), path("time-iavlTree-100.out"))
}

func BenchmarkIavlTreeBlock1000(b *testing.B) {
	runBlockBenchmark(b, newIavlTree, 1000, path("mem-iavlTree-1000.out"), path("time-iavlTree-1000.out"))
}

func BenchmarkIavlTreeBlock10000(b *testing.B) {
	runBlockBenchmark(b, newIavlTree, 10000, path("mem-iavlTree-10000.out"), path("time-iavlTree-10000.out"))
}

func BenchmarkIavlTreeBlock40000(b *testing.B) {
	runBlockBenchmark(b, newIavlTree, 40000, path("mem-iavlTree-40000.out"), path("time-iavlTree-40000.out"))
}
