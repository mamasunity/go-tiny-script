package benchmark

import (
	"sort"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	input := []string{"banana", "grape", "banana", "apple"}

	for i := 0; i < b.N; i++ {
		sort.Strings(input)
	}
}
