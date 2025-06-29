package day25_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day25"
)

func ExampleSolve() {
	day25.Solve("inputs/day25/sample-input.txt")
	// Output:
	// ANSWER: handshake encryption key: 14897079
}

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
