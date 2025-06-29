package day13_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day13"
)

func ExampleSolve() {
	day13.Solve("inputs/day13/sample-input.txt")
	// Output:
	// ANSWER1: product of earliest bus ID and wait time: 295
	// ANSWER2: earliest timestamp where all bus IDs depart at their respective offsets: 1068781
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
		day13.Solve("inputs/day13/input.txt")
	}
}
