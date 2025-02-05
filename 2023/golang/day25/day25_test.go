package day25_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day25"
)

func BenchmarkSolve(b *testing.B) {
	// b.ReportAllocs()

	oldStdout := os.Stdout
	null, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}
	os.Stdout = null

	defer func() {
		os.Stdout = oldStdout
		null.Close()
	}()

	b.ResetTimer()

	for range b.N {
		day25.Solve("inputs/day25/input.txt")
	}
}
