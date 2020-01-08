package benchs

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/dshulyak/seriesrw"
	"github.com/stretchr/testify/require"
)

func runBlockBenchmark(b *testing.B, initializer func(string) (testTree, error), size int, mempath, timepath string) {
	tmp, err := ioutil.TempDir("", "bench-block-")
	require.NoError(b, err)
	defer os.RemoveAll(tmp)

	tree, err := initializer(tmp)
	require.NoError(b, err)

	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)
	alloc := stats.Alloc
	total := 0

	memw, err := seriesrw.NewFileWriter(mempath, b.N*8)
	require.NoError(b, err)
	defer memw.Close()
	timew, err := seriesrw.NewFileWriter(timepath, b.N*8)
	require.NoError(b, err)
	defer timew.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		start := time.Now()
		for j := 0; j < size; j++ {
			key := make([]byte, 30)
			value := make([]byte, 100)
			rand.Read(key)
			rand.Read(value)
			require.NoError(b, tree.Write(key, value))
		}
		require.NoError(b, tree.Commit())
		runtime.ReadMemStats(&stats)
		total += size
		mem := stats.Alloc - alloc
		since := time.Since(start)
		fmt.Printf("commit took %v, memory %d mb, total %d\n", since, mem/1024/1024, total)
		require.NoError(b, memw.Write(mem))
		require.NoError(b, timew.Write(uint64(since)))
	}
}
