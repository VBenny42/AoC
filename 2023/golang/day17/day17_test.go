package day17_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day17"
)

func ExampleSolve() {
	day17.Solve("inputs/day17/sample-input.txt")
	// Output:
	// ANSWER1: least heat loss incurred: 102
	// ANSWER2: least heat loss incurred for ultra crucible: 94
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
		day17.Solve("inputs/day17/input.txt")
	}
}
