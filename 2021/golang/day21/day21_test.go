package day21_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day21"
)

func ExampleSolve() {
	day21.Solve("inputs/day21/sample-input.txt")
	// Output:
	// ANSWER1: losing play score times number of die rolls: 739785
	// ANSWER2: number of universes where the most likely winner wins: 444356092776315
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
		day21.Solve("inputs/day21/input.txt")
	}
}
