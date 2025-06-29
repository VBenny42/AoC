package day04_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day04"
)

func ExampleSolve() {
	day04.Solve("inputs/day04/sample-input.txt")
	// Output:
	// ANSWER1: number of valid passports: 2
	// ANSWER2: number of valid passports with valid values: 2
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
		day04.Solve("inputs/day04/input.txt")
	}
}
