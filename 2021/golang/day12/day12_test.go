package day12_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day12"
)

func ExampleSolve() {
	day12.Solve("inputs/day12/sample-input.txt")
	// Output:
	// ANSWER1: paths through cave system that visit small caves at least once: 19
	// ANSWER2: paths through cave system that visit one small cave twice: 103
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
		day12.Solve("inputs/day12/input.txt")
	}
}
