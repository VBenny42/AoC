package day20_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day20"
)

func ExampleSolve() {
	day20.Solve("inputs/day20/sample-input.txt")
	// Output:
	// ANSWER1: product of the IDs of the 4 corner tiles: 20899048083289
	// ANSWER2: water roughness after removing sea monsters: 273
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
		day20.Solve("inputs/day20/input.txt")
	}
}
