package day06_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day06"
)

func ExampleSolve() {
	day06.Solve("inputs/day06/sample-input.txt")
	// Output:
	// ANSWER1: number of lanternfish after 80 days: 5934
	// ANSWER2: number of lanternfish after 256 days: 26984457539
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
		day06.Solve("inputs/day06/input.txt")
	}
}
