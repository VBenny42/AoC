package day24_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day24"
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
		day24.Solve("inputs/day24/input.txt")
	}
}
