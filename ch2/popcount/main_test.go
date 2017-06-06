package popcount

import "testing"

func BenchmarkLoopPopCount(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		LoopPopCount(uint64(i))
	}
}


func BenchmarkPopCount(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}