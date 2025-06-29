package day02_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day02"
)

func ExampleSolve() {
	day02.Solve("inputs/day02/sample-input.txt")
	// Output:
	// ANSWER1: number of valid passwords: 2
	// ANSWER2: number of valid passwords according to new interpretation: 1
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
		day02.Solve("inputs/day02/input.txt")
	}
}
