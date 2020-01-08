package benchs

import (
	"testing"
)

func BenchmarkPatriciaBlock100(b *testing.B) {
	runBlockBenchmark(b, newPatricia, 100, path("mem-patricia-100.out"), path("time-patricia-100.out"))
}

func BenchmarkPatriciaBlock1000(b *testing.B) {
	runBlockBenchmark(b, newPatricia, 1000, path("mem-patricia-1000.out"), path("time-patricia-1000.out"))
}

func BenchmarkPatriciaBlock10000(b *testing.B) {
	runBlockBenchmark(b, newPatricia, 10000, path("mem-patricia-10000.out"), path("time-patricia-10000.out"))
}

func BenchmarkPatriciaBlock40000(b *testing.B) {
	runBlockBenchmark(b, newPatricia, 40000, path("mem-patricia-40000.out"), path("time-patricia-40000.out"))
}
